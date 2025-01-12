package commands

import (
	"github.com/nebojsa94/erigon/cmd/state/verify"
	"github.com/nebojsa94/erigon/turbo/debug"
	"github.com/spf13/cobra"
)

func init() {
	withDataDir(verifyTxLookupCmd)
	rootCmd.AddCommand(verifyTxLookupCmd)
}

var verifyTxLookupCmd = &cobra.Command{
	Use:   "verifyTxLookup",
	Short: "Generate tx lookup index",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := debug.SetupCobra(cmd, "verify_txlookup")
		return verify.ValidateTxLookups(chaindata, logger)
	},
}
