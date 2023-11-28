package statechange

import (
	"github.com/nebojsa94/erigon/cl/abstract"
	"github.com/nebojsa94/erigon/cl/clparams"
	"github.com/nebojsa94/erigon/cl/cltypes"
	"github.com/nebojsa94/erigon/cl/phase1/core/state"
	"github.com/nebojsa94/erigon/cl/utils"
)

// ProcessHistoricalRootsUpdate updates the historical root data structure by computing a new historical root batch when it is time to do so.
func ProcessHistoricalRootsUpdate(s abstract.BeaconState) error {
	nextEpoch := state.Epoch(s) + 1
	beaconConfig := s.BeaconConfig()
	blockRoots := s.BlockRoots()
	stateRoots := s.StateRoots()

	// Check if it's time to compute the historical root batch.
	if nextEpoch%(beaconConfig.SlotsPerHistoricalRoot/beaconConfig.SlotsPerEpoch) != 0 {
		return nil
	}

	// Compute historical root batch.
	blockRootsLeaf, err := blockRoots.HashSSZ()
	if err != nil {
		return err
	}
	stateRootsLeaf, err := stateRoots.HashSSZ()
	if err != nil {
		return err
	}

	// Add the historical summary or root to the s.
	if s.Version() >= clparams.CapellaVersion {
		s.AddHistoricalSummary(&cltypes.HistoricalSummary{
			BlockSummaryRoot: blockRootsLeaf,
			StateSummaryRoot: stateRootsLeaf,
		})
	} else {
		historicalRoot := utils.Sha256(blockRootsLeaf[:], stateRootsLeaf[:])
		s.AddHistoricalRoot(historicalRoot)
	}

	return nil
}
