package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/nebojsa94/erigon/cmd/rpcdaemon/cli"
	"github.com/nebojsa94/erigon/rpc"
	"github.com/nebojsa94/erigon/turbo/debug"
	"github.com/nebojsa94/erigon/turbo/jsonrpc"
	"github.com/spf13/cobra"
)

func main() {
	cmd, cfg := cli.RootCommand()
	rootCtx, rootCancel := common.RootContext()
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		logger := debug.SetupCobra(cmd, "sentry")
		db, backend, txPool, mining, stateCache, blockReader, engine, ff, agg, err := cli.RemoteServices(ctx, cfg, logger, rootCancel)
		if err != nil {
			if !errors.Is(err, context.Canceled) {
				logger.Error("Could not connect to DB", "err", err)
			}
			return nil
		}
		defer db.Close()
		defer engine.Close()

		apiList := jsonrpc.APIList(db, backend, txPool, mining, ff, stateCache, blockReader, agg, cfg, engine, logger)
		rpc.PreAllocateRPCMetricLabels(apiList)
		if err := cli.StartRpcServer(ctx, cfg, apiList, logger); err != nil {
			logger.Error(err.Error())
			return nil
		}

		return nil
	}

	if err := cmd.ExecuteContext(rootCtx); err != nil {
		fmt.Printf("ExecuteContext: %v\n", err)
		os.Exit(1)
	}
}
