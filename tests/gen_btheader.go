// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"math/big"

	libcommon "github.com/nebojsa94/erigon/erigon-lib/common"
	"github.com/nebojsa94/erigon/erigon-lib/common/hexutility"

	"github.com/nebojsa94/erigon/common/math"
	"github.com/nebojsa94/erigon/core/types"
)

var _ = (*btHeaderMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (b btHeader) MarshalJSON() ([]byte, error) {
	type btHeader struct {
		Bloom                 types.Bloom
		Coinbase              libcommon.Address
		MixHash               libcommon.Hash
		Nonce                 types.BlockNonce
		Number                *math.HexOrDecimal256
		Hash                  libcommon.Hash
		ParentHash            libcommon.Hash
		ReceiptTrie           libcommon.Hash
		StateRoot             libcommon.Hash
		TransactionsTrie      libcommon.Hash
		UncleHash             libcommon.Hash
		ExtraData             hexutility.Bytes
		Difficulty            *math.HexOrDecimal256
		GasLimit              math.HexOrDecimal64
		GasUsed               math.HexOrDecimal64
		Timestamp             math.HexOrDecimal64
		BaseFeePerGas         *math.HexOrDecimal256
		WithdrawalsRoot       *libcommon.Hash
		BlobGasUsed           *math.HexOrDecimal64
		ExcessBlobGas         *math.HexOrDecimal64
		ParentBeaconBlockRoot *libcommon.Hash
	}
	var enc btHeader
	enc.Bloom = b.Bloom
	enc.Coinbase = b.Coinbase
	enc.MixHash = b.MixHash
	enc.Nonce = b.Nonce
	enc.Number = (*math.HexOrDecimal256)(b.Number)
	enc.Hash = b.Hash
	enc.ParentHash = b.ParentHash
	enc.ReceiptTrie = b.ReceiptTrie
	enc.StateRoot = b.StateRoot
	enc.TransactionsTrie = b.TransactionsTrie
	enc.UncleHash = b.UncleHash
	enc.ExtraData = b.ExtraData
	enc.Difficulty = (*math.HexOrDecimal256)(b.Difficulty)
	enc.GasLimit = math.HexOrDecimal64(b.GasLimit)
	enc.GasUsed = math.HexOrDecimal64(b.GasUsed)
	enc.Timestamp = math.HexOrDecimal64(b.Timestamp)
	enc.BaseFeePerGas = (*math.HexOrDecimal256)(b.BaseFeePerGas)
	enc.WithdrawalsRoot = b.WithdrawalsRoot
	enc.BlobGasUsed = (*math.HexOrDecimal64)(b.BlobGasUsed)
	enc.ExcessBlobGas = (*math.HexOrDecimal64)(b.ExcessBlobGas)
	enc.ParentBeaconBlockRoot = b.ParentBeaconBlockRoot
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (b *btHeader) UnmarshalJSON(input []byte) error {
	type btHeader struct {
		Bloom                 *types.Bloom
		Coinbase              *libcommon.Address
		MixHash               *libcommon.Hash
		Nonce                 *types.BlockNonce
		Number                *math.HexOrDecimal256
		Hash                  *libcommon.Hash
		ParentHash            *libcommon.Hash
		ReceiptTrie           *libcommon.Hash
		StateRoot             *libcommon.Hash
		TransactionsTrie      *libcommon.Hash
		UncleHash             *libcommon.Hash
		ExtraData             *hexutility.Bytes
		Difficulty            *math.HexOrDecimal256
		GasLimit              *math.HexOrDecimal64
		GasUsed               *math.HexOrDecimal64
		Timestamp             *math.HexOrDecimal64
		BaseFeePerGas         *math.HexOrDecimal256
		WithdrawalsRoot       *libcommon.Hash
		BlobGasUsed           *math.HexOrDecimal64
		ExcessBlobGas         *math.HexOrDecimal64
		ParentBeaconBlockRoot *libcommon.Hash
	}
	var dec btHeader
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Bloom != nil {
		b.Bloom = *dec.Bloom
	}
	if dec.Coinbase != nil {
		b.Coinbase = *dec.Coinbase
	}
	if dec.MixHash != nil {
		b.MixHash = *dec.MixHash
	}
	if dec.Nonce != nil {
		b.Nonce = *dec.Nonce
	}
	if dec.Number != nil {
		b.Number = (*big.Int)(dec.Number)
	}
	if dec.Hash != nil {
		b.Hash = *dec.Hash
	}
	if dec.ParentHash != nil {
		b.ParentHash = *dec.ParentHash
	}
	if dec.ReceiptTrie != nil {
		b.ReceiptTrie = *dec.ReceiptTrie
	}
	if dec.StateRoot != nil {
		b.StateRoot = *dec.StateRoot
	}
	if dec.TransactionsTrie != nil {
		b.TransactionsTrie = *dec.TransactionsTrie
	}
	if dec.UncleHash != nil {
		b.UncleHash = *dec.UncleHash
	}
	if dec.ExtraData != nil {
		b.ExtraData = *dec.ExtraData
	}
	if dec.Difficulty != nil {
		b.Difficulty = (*big.Int)(dec.Difficulty)
	}
	if dec.GasLimit != nil {
		b.GasLimit = uint64(*dec.GasLimit)
	}
	if dec.GasUsed != nil {
		b.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.Timestamp != nil {
		b.Timestamp = uint64(*dec.Timestamp)
	}
	if dec.BaseFeePerGas != nil {
		b.BaseFeePerGas = (*big.Int)(dec.BaseFeePerGas)
	}
	if dec.WithdrawalsRoot != nil {
		b.WithdrawalsRoot = dec.WithdrawalsRoot
	}
	if dec.BlobGasUsed != nil {
		b.BlobGasUsed = (*uint64)(dec.BlobGasUsed)
	}
	if dec.ExcessBlobGas != nil {
		b.ExcessBlobGas = (*uint64)(dec.ExcessBlobGas)
	}
	if dec.ParentBeaconBlockRoot != nil {
		b.ParentBeaconBlockRoot = dec.ParentBeaconBlockRoot
	}
	return nil
}
