package generics

import (
	"sync/atomic"

	libcommon "github.com/nebojsa94/erigon-lib/common"
	"github.com/nebojsa94/erigon/core/types"
)

func Empty[T any]() (t T) {
	return
}

// BorMilestoneRewind is used as a flag/variable
// Flag: if equals 0, no rewind according to bor whitelisting service
// Variable: if not equals 0, rewind chain back to BorMilestoneRewind
var BorMilestoneRewind atomic.Pointer[uint64]

type Response struct {
	Headers []*types.Header
	Hashes  []libcommon.Hash
}
