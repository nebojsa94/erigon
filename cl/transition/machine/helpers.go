package machine

import (
	"github.com/nebojsa94/erigon/erigon-lib/common"
	"github.com/nebojsa94/erigon/cl/abstract"
	"github.com/nebojsa94/erigon/cl/cltypes"
	"github.com/nebojsa94/erigon/cl/phase1/core/state"
)

func executionEnabled(s abstract.BeaconState, payload *cltypes.Eth1Block) bool {
	return (!state.IsMergeTransitionComplete(s) && payload.BlockHash != common.Hash{}) || state.IsMergeTransitionComplete(s)
}
