// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package native

import (
	"encoding/json"
	"github.com/nebojsa94/erigon-lib/common/hexutil"
	"math/big"

	libcommon "github.com/nebojsa94/erigon-lib/common"
	"github.com/nebojsa94/erigon-lib/common/hexutility"

	"github.com/nebojsa94/erigon/core/vm"
)

var _ = (*callFrameMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (c callFrame) MarshalJSON() ([]byte, error) {
	type callFrame0 struct {
		Type       vm.OpCode         `json:"-"`
		From       libcommon.Address `json:"from"`
		Gas        hexutil.Uint64    `json:"gas"`
		GasUsed    hexutil.Uint64    `json:"gasUsed"`
		To         libcommon.Address `json:"to,omitempty" rlp:"optional"`
		Input      hexutility.Bytes  `json:"input" rlp:"optional"`
		Output     hexutility.Bytes  `json:"output,omitempty" rlp:"optional"`
		Error      string            `json:"error,omitempty" rlp:"optional"`
		Revertal   string            `json:"revertReason,omitempty"`
		Calls      []callFrame       `json:"calls,omitempty" rlp:"optional"`
		Logs       []callLog         `json:"logs,omitempty" rlp:"optional"`
		Value      *hexutil.Big      `json:"value,omitempty" rlp:"optional"`
		TypeString string            `json:"type"`
	}
	var enc callFrame0
	enc.Type = c.Type
	enc.From = c.From
	enc.Gas = hexutil.Uint64(c.Gas)
	enc.GasUsed = hexutil.Uint64(c.GasUsed)
	enc.To = c.To
	enc.Input = c.Input
	enc.Output = c.Output
	enc.Error = c.Error
	enc.Revertal = c.Revertal
	enc.Calls = c.Calls
	enc.Logs = c.Logs
	enc.Value = (*hexutil.Big)(c.Value)
	enc.TypeString = c.TypeString()
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (c *callFrame) UnmarshalJSON(input []byte) error {
	type callFrame0 struct {
		Type     *vm.OpCode         `json:"-"`
		From     *libcommon.Address `json:"from"`
		Gas      *hexutil.Uint64    `json:"gas"`
		GasUsed  *hexutil.Uint64    `json:"gasUsed"`
		To       *libcommon.Address `json:"to,omitempty" rlp:"optional"`
		Input    *hexutility.Bytes  `json:"input" rlp:"optional"`
		Output   *hexutility.Bytes  `json:"output,omitempty" rlp:"optional"`
		Error    *string            `json:"error,omitempty" rlp:"optional"`
		Revertal *string            `json:"revertReason,omitempty"`
		Calls    []callFrame        `json:"calls,omitempty" rlp:"optional"`
		Logs     []callLog          `json:"logs,omitempty" rlp:"optional"`
		Value    *hexutil.Big       `json:"value,omitempty" rlp:"optional"`
	}
	var dec callFrame0
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Type != nil {
		c.Type = *dec.Type
	}
	if dec.From != nil {
		c.From = *dec.From
	}
	if dec.Gas != nil {
		c.Gas = uint64(*dec.Gas)
	}
	if dec.GasUsed != nil {
		c.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.To != nil {
		c.To = *dec.To
	}
	if dec.Input != nil {
		c.Input = *dec.Input
	}
	if dec.Output != nil {
		c.Output = *dec.Output
	}
	if dec.Error != nil {
		c.Error = *dec.Error
	}
	if dec.Revertal != nil {
		c.Revertal = *dec.Revertal
	}
	if dec.Calls != nil {
		c.Calls = dec.Calls
	}
	if dec.Logs != nil {
		c.Logs = dec.Logs
	}
	if dec.Value != nil {
		c.Value = (*big.Int)(dec.Value)
	}
	return nil
}
