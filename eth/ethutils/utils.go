package ethutils

import (
	libcommon "github.com/nebojsa94/erigon/erigon-lib/common"
	"github.com/ledgerwatch/log/v3"

	"github.com/nebojsa94/erigon/consensus"
	"github.com/nebojsa94/erigon/core/types"
)

// IsLocalBlock checks whether the specified block is mined
// by local miner accounts.
//
// We regard two types of accounts as local miner account: etherbase
// and accounts specified via `txpool.locals` flag.
func IsLocalBlock(engine consensus.Engine, etherbase libcommon.Address, txPoolLocals []libcommon.Address, header *types.Header) bool {
	author, err := engine.Author(header)
	if err != nil {
		log.Warn("Failed to retrieve block author", "number", header.Number, "header_hash", header.Hash(), "err", err)
		return false
	}
	// Check whether the given address is etherbase.
	if author == etherbase {
		return true
	}
	// Check whether the given address is specified by `txpool.local`
	// CLI flag.
	for _, account := range txPoolLocals {
		if account == author {
			return true
		}
	}
	return false
}
