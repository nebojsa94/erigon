package main

import (
	"fmt"
	"os"

	"github.com/nebojsa94/erigon-lib/common"
	"github.com/nebojsa94/erigon/cmd/integration/commands"
)

func main() {
	rootCmd := commands.RootCommand()
	ctx, _ := common.RootContext()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
