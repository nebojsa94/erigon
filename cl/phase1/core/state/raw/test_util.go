package raw

import (
	_ "embed"

	"github.com/nebojsa94/erigon/cl/clparams"
	"github.com/nebojsa94/erigon/cl/utils"
)

//go:embed testdata/state.ssz_snappy
var denebState []byte

func GetTestState() *BeaconState {
	state := New(&clparams.MainnetBeaconConfig)
	utils.DecodeSSZSnappy(state, denebState, int(clparams.DenebVersion))
	return state

}
