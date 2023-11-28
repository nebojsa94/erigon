package tracers

import (
	"encoding/json"

	"github.com/nebojsa94/erigon/erigon-lib/common/hexutil"
	"github.com/nebojsa94/erigon/eth/tracers/logger"
	"github.com/nebojsa94/erigon/turbo/adapter/ethapi"
)

// TraceConfig holds extra parameters to trace functions.
type TraceConfig struct {
	*logger.LogConfig
	Tracer         *string
	TracerConfig   *json.RawMessage
	Timeout        *string
	Reexec         *uint64
	NoRefunds      *bool // Turns off gas refunds when tracing
	StateOverrides *ethapi.StateOverrides

	BorTraceEnabled *bool
	BorTx           *bool
	TxIndex         *hexutil.Uint
}
