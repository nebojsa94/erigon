package jsonrpc

import (
	"context"
	"fmt"
	"testing"

	libcommon "github.com/nebojsa94/erigon/erigon-lib/common"
	"github.com/nebojsa94/erigon/erigon-lib/direct"
	"github.com/nebojsa94/erigon/erigon-lib/gointerfaces/sentry"
	"github.com/stretchr/testify/require"

	"github.com/nebojsa94/erigon/cmd/rpcdaemon/rpcservices"
	"github.com/nebojsa94/erigon/core"
	"github.com/nebojsa94/erigon/eth/protocols/eth"
	"github.com/nebojsa94/erigon/ethdb/privateapi"
	"github.com/nebojsa94/erigon/rlp"
	"github.com/nebojsa94/erigon/turbo/builder"
	"github.com/nebojsa94/erigon/turbo/rpchelper"
	"github.com/nebojsa94/erigon/turbo/stages"
	"github.com/nebojsa94/erigon/turbo/stages/mock"
	"github.com/ledgerwatch/log/v3"
)

func TestEthSubscribe(t *testing.T) {
	m, require := mock.Mock(t), require.New(t)
	chain, err := core.GenerateChain(m.ChainConfig, m.Genesis, m.Engine, m.DB, 7, func(i int, b *core.BlockGen) {
		b.SetCoinbase(libcommon.Address{1})
	})
	require.NoError(err)

	b, err := rlp.EncodeToBytes(&eth.BlockHeadersPacket66{
		RequestId:          1,
		BlockHeadersPacket: chain.Headers,
	})
	require.NoError(err)

	m.ReceiveWg.Add(1)
	for _, err = range m.Send(&sentry.InboundMessage{Id: sentry.MessageId_BLOCK_HEADERS_66, Data: b, PeerId: m.PeerId}) {
		require.NoError(err)
	}
	m.ReceiveWg.Wait() // Wait for all messages to be processed before we proceeed

	ctx := context.Background()
	logger := log.New()
	backendServer := privateapi.NewEthBackendServer(ctx, nil, m.DB, m.Notifications.Events, m.BlockReader, logger, builder.NewLatestBlockBuiltStore())
	backendClient := direct.NewEthBackendClientDirect(backendServer)
	backend := rpcservices.NewRemoteBackend(backendClient, m.DB, m.BlockReader)
	ff := rpchelper.New(ctx, backend, nil, nil, func() {}, m.Log)

	newHeads, id := ff.SubscribeNewHeads(16)
	defer ff.UnsubscribeHeads(id)

	initialCycle := mock.MockInsertAsInitialCycle
	highestSeenHeader := chain.TopBlock.NumberU64()

	hook := stages.NewHook(m.Ctx, m.DB, m.Notifications, m.Sync, m.BlockReader, m.ChainConfig, m.Log, m.UpdateHead)
	if err := stages.StageLoopIteration(m.Ctx, m.DB, nil, m.Sync, initialCycle, logger, m.BlockReader, hook, false); err != nil {
		t.Fatal(err)
	}

	for i := uint64(1); i <= highestSeenHeader; i++ {
		header := <-newHeads
		fmt.Printf("Got header %d\n", header.Number.Uint64())
		require.Equal(i, header.Number.Uint64())
	}
}
