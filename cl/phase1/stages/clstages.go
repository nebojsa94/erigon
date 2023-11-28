package stages

import (
	"context"
	"errors"
	"runtime"
	"time"

	"github.com/nebojsa94/erigon/erigon-lib/common"
	"github.com/nebojsa94/erigon/erigon-lib/common/dbg"
	"github.com/nebojsa94/erigon/erigon-lib/kv"
	"github.com/nebojsa94/erigon/cl/antiquary"
	"github.com/nebojsa94/erigon/cl/beacon/synced_data"
	"github.com/nebojsa94/erigon/cl/clparams"
	"github.com/nebojsa94/erigon/cl/clstages"
	"github.com/nebojsa94/erigon/cl/cltypes"
	"github.com/nebojsa94/erigon/cl/persistence"
	"github.com/nebojsa94/erigon/cl/persistence/beacon_indicies"
	"github.com/nebojsa94/erigon/cl/persistence/db_config"
	state_accessors "github.com/nebojsa94/erigon/cl/persistence/state"
	"github.com/nebojsa94/erigon/cl/phase1/core/state"
	"github.com/nebojsa94/erigon/cl/phase1/execution_client"
	"github.com/nebojsa94/erigon/cl/phase1/forkchoice"
	"github.com/nebojsa94/erigon/core/types"
	"github.com/nebojsa94/erigon/turbo/snapshotsync/freezeblocks"

	network2 "github.com/nebojsa94/erigon/cl/phase1/network"
	"github.com/nebojsa94/erigon/cl/rpc"
	"github.com/nebojsa94/erigon/cl/sentinel/peers"
	"github.com/nebojsa94/erigon/cl/utils"
	"github.com/ledgerwatch/log/v3"
)

type Cfg struct {
	rpc             *rpc.BeaconRpcP2P
	genesisCfg      *clparams.GenesisConfig
	beaconCfg       *clparams.BeaconChainConfig
	executionClient execution_client.ExecutionEngine
	state           *state.CachingBeaconState
	gossipManager   *network2.GossipManager
	forkChoice      *forkchoice.ForkChoiceStore
	beaconDB        persistence.BeaconChainDatabase
	indiciesDB      kv.RwDB
	tmpdir          string
	dbConfig        db_config.DatabaseConfiguration
	sn              *freezeblocks.CaplinSnapshots
	antiquary       *antiquary.Antiquary
	syncedData      *synced_data.SyncedDataManager

	hasDownloaded, backfilling bool
}

type Args struct {
	peers uint64

	targetEpoch, seenEpoch uint64
	targetSlot, seenSlot   uint64

	hasDownloaded bool
}

func ClStagesCfg(
	rpc *rpc.BeaconRpcP2P,
	antiquary *antiquary.Antiquary,
	genesisCfg *clparams.GenesisConfig,
	beaconCfg *clparams.BeaconChainConfig,
	state *state.CachingBeaconState,
	executionClient execution_client.ExecutionEngine,
	gossipManager *network2.GossipManager,
	forkChoice *forkchoice.ForkChoiceStore,
	beaconDB persistence.BeaconChainDatabase,
	indiciesDB kv.RwDB,
	sn *freezeblocks.CaplinSnapshots,
	tmpdir string,
	dbConfig db_config.DatabaseConfiguration,
	backfilling bool,
	syncedData *synced_data.SyncedDataManager,
) *Cfg {
	return &Cfg{
		rpc:             rpc,
		antiquary:       antiquary,
		genesisCfg:      genesisCfg,
		beaconCfg:       beaconCfg,
		state:           state,
		executionClient: executionClient,
		gossipManager:   gossipManager,
		forkChoice:      forkChoice,
		tmpdir:          tmpdir,
		beaconDB:        beaconDB,
		indiciesDB:      indiciesDB,
		dbConfig:        dbConfig,
		sn:              sn,
		backfilling:     backfilling,
		syncedData:      syncedData,
	}
}

type StageName = string

const (
	WaitForPeers             StageName = "WaitForPeers"
	CatchUpEpochs            StageName = "CatchUpEpochs"
	CatchUpBlocks            StageName = "CatchUpBlocks"
	ForkChoice               StageName = "ForkChoice"
	ListenForForks           StageName = "ListenForForks"
	CleanupAndPruning        StageName = "CleanupAndPruning"
	SleepForSlot             StageName = "SleepForSlot"
	DownloadHistoricalBlocks StageName = "DownloadHistoricalBlocks"
)

const (
	minPeersForDownload = uint64(4)
)

func MetaCatchingUp(args Args) StageName {
	if args.peers < minPeersForDownload {
		return WaitForPeers
	}
	if !args.hasDownloaded {
		return DownloadHistoricalBlocks
	}
	if args.seenEpoch < args.targetEpoch {
		return CatchUpEpochs
	}
	if args.seenSlot < args.targetSlot {
		return CatchUpBlocks
	}

	return ""
}

/*

this graph describes the state transitions for cl

digraph {
    compound=true;
    subgraph cluster_0 {
        label="syncing";
        WaitForPeers;
        CatchUpBlocks;
        CatchUpEpochs;
    }

    subgraph cluster_3 {
        label="if behind (transition function)"
        MetaCatchingUp;
    }

    subgraph cluster_1 {
        label="head";
        ForkChoice; CleanupAndPruning; ListenForForks; SleepForSlot;
    }

    MetaCatchingUp -> WaitForPeers
    MetaCatchingUp -> CatchUpEpochs
    MetaCatchingUp -> CatchUpBlocks

    WaitForPeers -> MetaCatchingUp[lhead=cluster_3]
    CatchUpEpochs -> MetaCatchingUp[lhead=cluster_3]
    CatchUpBlocks -> MetaCatchingUp[lhead=cluster_3]
    CleanupAndPruning -> MetaCatchingUp[lhead=cluster_3]
    ListenForForks -> MetaCatchingUp[lhead=cluster_3]
    ForkChoice -> MetaCatchingUp[lhead=cluster_3]

    CatchUpBlocks -> ForkChoice
    ForkChoice -> ListenForForks

    SleepForSlot -> WaitForPeers

    ListenForForks -> ForkChoice
    ListenForForks -> SleepForSlot
    ListenForForks -> CleanupAndPruning
    CleanupAndPruning -> SleepForSlot
}
*/

// ConsensusClStages creates a stage loop container to be used to run caplin
func ConsensusClStages(ctx context.Context,
	cfg *Cfg,
) *clstages.StageGraph[*Cfg, Args] {

	rpcSource := persistence.NewBeaconRpcSource(cfg.rpc)
	gossipSource := persistence.NewGossipSource(ctx, cfg.gossipManager)
	processBlock := func(tx kv.RwTx, block *cltypes.SignedBeaconBlock, newPayload, fullValidation bool) error {
		if err := cfg.forkChoice.OnBlock(block, newPayload, fullValidation); err != nil {
			log.Warn("fail to process block", "reason", err, "slot", block.Block.Slot)
			return err
		}
		if err := beacon_indicies.WriteHighestFinalized(tx, cfg.forkChoice.FinalizedSlot()); err != nil {
			return err
		}
		// Write block to database optimistically if we are very behind.
		return cfg.beaconDB.WriteBlock(ctx, tx, block, false)
	}

	// TODO: this is an ugly hack, but it works! Basically, we want shared state in the clstages.
	// Probably the correct long term solution is to create a third generic parameter that defines shared state
	// but for now, all it would have are the two gossip sources and the forkChoicesSinceReorg, so i don't think its worth it (yet).
	shouldForkChoiceSinceReorg := false

	// clstages run in a single thread - so we don't need to worry about any synchronization.
	return &clstages.StageGraph[*Cfg, Args]{
		// the ArgsFunc is run after every stage. It is passed into the transition function, and the same args are passed into the next stage.
		ArgsFunc: func(ctx context.Context, cfg *Cfg) (args Args) {
			var err error
			args.peers, err = cfg.rpc.Peers()
			if err != nil {
				log.Error("failed to get sentinel peer count", "err", err)
				args.peers = 0
			}
			args.hasDownloaded = cfg.hasDownloaded
			args.seenSlot = cfg.forkChoice.HighestSeen()
			args.seenEpoch = args.seenSlot / cfg.beaconCfg.SlotsPerEpoch
			args.targetSlot = utils.GetCurrentSlot(cfg.genesisCfg.GenesisTime, cfg.beaconCfg.SecondsPerSlot)
			// Note that the target epoch is always one behind. this is because we are always behind in the current epoch, so it would not be very useful
			args.targetEpoch = utils.GetCurrentEpoch(cfg.genesisCfg.GenesisTime, cfg.beaconCfg.SecondsPerSlot, cfg.beaconCfg.SlotsPerEpoch) - 1
			return
		},
		Stages: map[string]clstages.Stage[*Cfg, Args]{
			WaitForPeers: {
				Description: `wait for enough peers. This is also a safe stage to go to when unsure of what stage to use`,
				TransitionFunc: func(cfg *Cfg, args Args, err error) string {
					if x := MetaCatchingUp(args); x != "" {
						return x
					}
					return CatchUpBlocks
				},
				ActionFunc: func(ctx context.Context, logger log.Logger, cfg *Cfg, args Args) error {
					peersCount, err := cfg.rpc.Peers()
					if err != nil {
						return nil
					}
					waitWhenNotEnoughPeers := 3 * time.Second
					for {
						if peersCount >= minPeersForDownload {
							break
						}
						select {
						case <-ctx.Done():
							return ctx.Err()
						default:
						}
						logger.Info("[Caplin] Waiting For Peers", "have", peersCount, "needed", minPeersForDownload, "retryIn", waitWhenNotEnoughPeers)
						time.Sleep(waitWhenNotEnoughPeers)
						peersCount, err = cfg.rpc.Peers()
						if err != nil {
							peersCount = 0
						}
					}
					return nil
				},
			},
			DownloadHistoricalBlocks: {
				Description: "Download historical blocks",
				TransitionFunc: func(cfg *Cfg, args Args, err error) string {
					if x := MetaCatchingUp(args); x != "" {
						return x
					}
					return CatchUpBlocks
				},
				ActionFunc: func(ctx context.Context, logger log.Logger, cfg *Cfg, args Args) error {
					cfg.hasDownloaded = true
					startingRoot, err := cfg.state.BlockRoot()
					if err != nil {
						return err
					}
					// This stage is special so use context.Background() TODO(Giulio2002): make the context be passed in
					startingSlot := cfg.state.LatestBlockHeader().Slot
					downloader := network2.NewBackwardBeaconDownloader(context.Background(), cfg.rpc, cfg.indiciesDB)

					if err := SpawnStageHistoryDownload(StageHistoryReconstruction(downloader, cfg.antiquary, cfg.sn, cfg.beaconDB, cfg.indiciesDB, cfg.executionClient, cfg.genesisCfg, cfg.beaconCfg, cfg.backfilling, false, startingRoot, startingSlot, cfg.tmpdir, logger), context.Background(), logger); err != nil {
						cfg.hasDownloaded = false
						return err
					}
					return nil
				},
			},
			CatchUpEpochs: {
				Description: `if we are 1 or more epochs behind, we download in parallel by epoch`,
				TransitionFunc: func(cfg *Cfg, args Args, err error) string {
					if x := MetaCatchingUp(args); x != "" {
						return x
					}
					return CatchUpBlocks
				},
				ActionFunc: func(ctx context.Context, logger log.Logger, cfg *Cfg, args Args) error {
					logger.Info("[Caplin] Downloading epochs from reqresp", "from", args.seenEpoch, "to", args.targetEpoch)
					currentEpoch := args.seenEpoch
					blockBatch := []*types.Block{}
					shouldInsert := cfg.executionClient != nil && cfg.executionClient.SupportInsertion()
					tx, err := cfg.indiciesDB.BeginRw(ctx)
					if err != nil {
						return err
					}
					defer tx.Rollback()
				MainLoop:
					for currentEpoch <= args.targetEpoch+1 {
						startBlock := currentEpoch * cfg.beaconCfg.SlotsPerEpoch
						blocks, err := rpcSource.GetRange(ctx, tx, startBlock, cfg.beaconCfg.SlotsPerEpoch)
						if err != nil {
							return err
						}
						// If we got an empty packet ban the peer
						if len(blocks.Data) == 0 {
							cfg.rpc.BanPeer(blocks.Peer)
							log.Debug("no data received from peer in epoch download")
							continue MainLoop
						}

						logger.Info("[Caplin] Epoch downloaded", "epoch", currentEpoch)
						for _, block := range blocks.Data {

							if shouldInsert && block.Version() >= clparams.BellatrixVersion {
								executionPayload := block.Block.Body.ExecutionPayload
								body := executionPayload.Body()
								txs, err := types.DecodeTransactions(body.Transactions)
								if err != nil {
									log.Warn("bad blocks segment received", "err", err)
									cfg.rpc.BanPeer(blocks.Peer)
									currentEpoch = utils.Max64(args.seenEpoch, currentEpoch-1)
									continue MainLoop
								}
								header, err := executionPayload.RlpHeader()
								if err != nil {
									log.Warn("bad blocks segment received", "err", err)
									cfg.rpc.BanPeer(blocks.Peer)
									currentEpoch = utils.Max64(args.seenEpoch, currentEpoch-1)
									continue MainLoop
								}
								blockBatch = append(blockBatch, types.NewBlockFromStorage(executionPayload.BlockHash, header, txs, nil, body.Withdrawals))
							}
							if err := processBlock(tx, block, false, true); err != nil {
								log.Warn("bad blocks segment received", "err", err)
								cfg.rpc.BanPeer(blocks.Peer)
								currentEpoch = utils.Max64(args.seenEpoch, currentEpoch-1)
								continue MainLoop
							}
						}
						if len(blockBatch) > 0 {
							if err := cfg.executionClient.InsertBlocks(blockBatch); err != nil {
								log.Warn("bad blocks segment received", "err", err)
								currentEpoch = utils.Max64(args.seenEpoch, currentEpoch-1)
								blockBatch = blockBatch[:0]
								continue MainLoop
							}
							blockBatch = blockBatch[:0]
						}
						currentEpoch++
					}
					return tx.Commit()
				},
			},
			CatchUpBlocks: {
				Description: `if we are within the epoch but not at head, we run catchupblocks`,
				TransitionFunc: func(cfg *Cfg, args Args, err error) string {
					if x := MetaCatchingUp(args); x != "" {
						return x
					}
					return ForkChoice
				},
				ActionFunc: func(ctx context.Context, logger log.Logger, cfg *Cfg, args Args) error {
					totalRequest := args.targetSlot - args.seenSlot
					logger.Debug("waiting for blocks...",
						"seenSlot", args.seenSlot,
						"targetSlot", args.targetSlot,
						"requestedSlots", totalRequest,
					)
					respCh := make(chan *peers.PeeredObject[[]*cltypes.SignedBeaconBlock])
					errCh := make(chan error)
					sources := []persistence.BlockSource{gossipSource}

					// if we are more than one block behind, we request the rpc source as well
					if totalRequest > 2 {
						sources = append(sources, rpcSource)
					}
					// 15 seconds is a good timeout for this
					ctx, cn := context.WithTimeout(ctx, 15*time.Second)
					defer cn()

					tx, err := cfg.indiciesDB.BeginRw(ctx)
					if err != nil {
						return err
					}
					defer tx.Rollback()
					// we go ask all the sources and see who gets back to us first. whoever does is the winner!!
					for _, v := range sources {
						sourceFunc := v.GetRange
						go func() {
							blocks, err := sourceFunc(ctx, tx, args.seenSlot+1, totalRequest)
							if err != nil {
								errCh <- err
								return
							}
							respCh <- blocks
						}()
					}
					logTimer := time.NewTicker(30 * time.Second)
					defer logTimer.Stop()
					select {
					case err := <-errCh:
						return err
					case blocks := <-respCh:
						for _, block := range blocks.Data {
							if err := processBlock(tx, block, true, true); err != nil {
								return err
							}
						}
					case <-logTimer.C:
						logger.Info("[Caplin] Progress", "progress", cfg.forkChoice.HighestSeen(), "from", args.seenEpoch, "to", args.targetSlot)
					}
					return tx.Commit()
				},
			},
			ForkChoice: {
				Description: `fork choice stage. We will send all fork choise things here
			also, we will wait up to delay seconds to deal with attestations + side forks`,
				TransitionFunc: func(cfg *Cfg, args Args, err error) string {
					if x := MetaCatchingUp(args); x != "" {
						return x
					}
					return ListenForForks
				},
				ActionFunc: func(ctx context.Context, logger log.Logger, cfg *Cfg, args Args) error {

					// TODO: we need to get the last run block in order to process attestations here
					////////block.Block.Body.Attestations.Range(func(idx int, a *solid.Attestation, total int) bool {
					////////	if err = g.forkChoice.OnAttestation(a, true); err != nil {
					////////		return false
					////////	}
					////////	return true
					////////})
					////////if err != nil {
					////////	return err
					////////}

					// Now check the head
					headRoot, headSlot, err := cfg.forkChoice.GetHead()
					if err != nil {
						return err
					}

					// Do forkchoice if possible
					if cfg.forkChoice.Engine() != nil {
						finalizedCheckpoint := cfg.forkChoice.FinalizedCheckpoint()
						logger.Debug("Caplin is sending forkchoice")
						// Run forkchoice
						if err := cfg.forkChoice.Engine().ForkChoiceUpdate(
							cfg.forkChoice.GetEth1Hash(finalizedCheckpoint.BlockRoot()),
							cfg.forkChoice.GetEth1Hash(headRoot),
						); err != nil {
							logger.Warn("Could not set forkchoice", "err", err)
							return err
						}
					}
					tx, err := cfg.indiciesDB.BeginRw(ctx)
					if err != nil {
						return err
					}
					defer tx.Rollback()
					// Fix canonical chain in the indexed datatabase.
					if err := beacon_indicies.TruncateCanonicalChain(ctx, tx, headSlot); err != nil {
						return err
					}

					currentRoot := headRoot
					currentSlot := headSlot
					currentCanonical, err := beacon_indicies.ReadCanonicalBlockRoot(tx, currentSlot)
					if err != nil {
						return err
					}

					for currentRoot != currentCanonical {
						var newFoundSlot *uint64
						if err := beacon_indicies.MarkRootCanonical(ctx, tx, currentSlot, currentRoot); err != nil {
							return err
						}
						if currentRoot, err = beacon_indicies.ReadParentBlockRoot(ctx, tx, currentRoot); err != nil {
							return err
						}
						if newFoundSlot, err = beacon_indicies.ReadBlockSlotByBlockRoot(tx, currentRoot); err != nil {
							return err
						}
						if newFoundSlot == nil {
							break
						}
						currentSlot = *newFoundSlot
						currentCanonical, err = beacon_indicies.ReadCanonicalBlockRoot(tx, currentSlot)
						if err != nil {
							return err
						}
					}

					// Increment validator set
					headState, err := cfg.forkChoice.GetFullState(headRoot, false)
					if err != nil {
						return err
					}
					if err := cfg.syncedData.OnHeadState(headState); err != nil {
						return err
					}
					start := time.Now()
					preverifiedValidators := cfg.forkChoice.PreverifiedValidator(headState.FinalizedCheckpoint().BlockRoot())
					if err := state_accessors.IncrementPublicKeyTable(tx, headState, preverifiedValidators); err != nil {
						return err
					}
					log.Debug("Incremented state history", "elapsed", time.Since(start), "preverifiedValidators", preverifiedValidators)

					var m runtime.MemStats
					dbg.ReadMemStats(&m)
					logger.Debug("Imported chain segment",
						"hash", headRoot, "slot", headSlot,
						"alloc", common.ByteCount(m.Alloc),
						"sys", common.ByteCount(m.Sys))
					return tx.Commit()
				},
			},
			ListenForForks: {
				TransitionFunc: func(cfg *Cfg, args Args, err error) string {
					defer func() {
						shouldForkChoiceSinceReorg = false
					}()
					if x := MetaCatchingUp(args); x != "" {
						return x
					}
					if shouldForkChoiceSinceReorg {
						return ForkChoice
					}
					return CleanupAndPruning

				},
				ActionFunc: func(ctx context.Context, logger log.Logger, cfg *Cfg, args Args) error {
					slotTime := utils.GetSlotTime(cfg.genesisCfg.GenesisTime, cfg.beaconCfg.SecondsPerSlot, args.targetSlot).Add(
						time.Duration(cfg.beaconCfg.SecondsPerSlot) * (time.Second / 3),
					)
					waitDur := slotTime.Sub(time.Now())
					ctx, cn := context.WithTimeout(ctx, waitDur)
					defer cn()
					tx, err := cfg.indiciesDB.BeginRw(ctx)
					if err != nil {
						return err
					}
					defer tx.Rollback()
					// try to get the current block
					blocks, err := gossipSource.GetRange(ctx, tx, args.seenSlot, 1)
					if err != nil {
						if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
							return nil
						}
						return err
					}

					for _, block := range blocks.Data {
						err := processBlock(tx, block, true, true)
						if err != nil {
							// its okay if block processing fails
							logger.Warn("extra block failed validation", "err", err)
							return nil
						}
						shouldForkChoiceSinceReorg = true
						logger.Debug("extra block received", "slot", args.seenSlot)
					}
					return tx.Commit()
				},
			},
			CleanupAndPruning: {
				Description: `cleanup and pruning is done here`,
				TransitionFunc: func(cfg *Cfg, args Args, err error) string {
					if x := MetaCatchingUp(args); x != "" {
						return x
					}
					return SleepForSlot
				},
				ActionFunc: func(ctx context.Context, logger log.Logger, cfg *Cfg, args Args) error {
					tx, err := cfg.indiciesDB.BeginRw(ctx)
					if err != nil {
						return err
					}
					defer tx.Rollback()
					// clean up some old ranges
					err = gossipSource.PurgeRange(ctx, tx, 1, args.seenSlot-cfg.beaconCfg.SlotsPerEpoch*16)
					if err != nil {
						return err
					}
					// TODO(Giulio2002): schedule snapshots retirement if needed.
					if !cfg.backfilling {
						if err := cfg.beaconDB.PurgeRange(ctx, tx, 1, cfg.forkChoice.HighestSeen()-100_000); err != nil {
							return err
						}
						if err := beacon_indicies.PruneBlockRoots(ctx, tx, 0, cfg.forkChoice.HighestSeen()-100_000); err != nil {
							return err
						}
					}

					return tx.Commit()
				},
			},
			SleepForSlot: {
				Description: `sleep until the next slot`,
				TransitionFunc: func(cfg *Cfg, args Args, err error) string {
					return WaitForPeers
				},
				ActionFunc: func(ctx context.Context, logger log.Logger, cfg *Cfg, args Args) error {
					nextSlot := args.seenSlot + 1
					nextSlotTime := utils.GetSlotTime(cfg.genesisCfg.GenesisTime, cfg.beaconCfg.SecondsPerSlot, nextSlot)
					nextSlotDur := nextSlotTime.Sub(time.Now())
					logger.Debug("sleeping until next slot", "slot", nextSlot, "time", nextSlotTime, "dur", nextSlotDur)
					time.Sleep(nextSlotDur)
					return nil
				},
			},
		},
	}
}
