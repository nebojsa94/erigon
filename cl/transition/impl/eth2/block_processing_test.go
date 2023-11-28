package eth2_test

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nebojsa94/erigon/cl/clparams"
	"github.com/nebojsa94/erigon/cl/cltypes"
	"github.com/nebojsa94/erigon/cl/phase1/core/state"
	"github.com/nebojsa94/erigon/cl/transition"
	"github.com/nebojsa94/erigon/cl/utils"
)

//go:embed statechange/test_data/block_processing/capella_block.ssz_snappy
var capellaBlock []byte

//go:embed statechange/test_data/block_processing/capella_state.ssz_snappy
var capellaState []byte

func TestBlockProcessing(t *testing.T) {
	s := state.New(&clparams.MainnetBeaconConfig)
	require.NoError(t, utils.DecodeSSZSnappy(s, capellaState, int(clparams.CapellaVersion)))
	block := cltypes.NewSignedBeaconBlock(&clparams.MainnetBeaconConfig)
	require.NoError(t, utils.DecodeSSZSnappy(block, capellaBlock, int(clparams.CapellaVersion)))
	require.NoError(t, transition.TransitionState(s, block, true)) // All checks already made in transition state
}
