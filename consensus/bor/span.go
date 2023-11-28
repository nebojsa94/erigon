package bor

import (
	libcommon "github.com/nebojsa94/erigon/erigon-lib/common"
	"github.com/nebojsa94/erigon/consensus"
	"github.com/nebojsa94/erigon/consensus/bor/heimdall/span"
	"github.com/nebojsa94/erigon/consensus/bor/valset"
)

//go:generate mockgen -destination=./span_mock.go -package=bor . Spanner
type Spanner interface {
	GetCurrentSpan(syscall consensus.SystemCall) (*span.Span, error)
	GetCurrentValidators(spanId uint64, signer libcommon.Address, chain consensus.ChainHeaderReader) ([]*valset.Validator, error)
	GetCurrentProducers(spanId uint64, signer libcommon.Address, chain consensus.ChainHeaderReader) ([]*valset.Validator, error)
	CommitSpan(heimdallSpan span.HeimdallSpan, syscall consensus.SystemCall) error
}
