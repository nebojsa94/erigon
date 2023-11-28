// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/nebojsa94/erigon/erigon-lib/chain"
	"github.com/nebojsa94/erigon/erigon-lib/common"
	"github.com/nebojsa94/erigon/erigon-lib/common/hexutility"
	common0 "github.com/nebojsa94/erigon/common"
	"github.com/nebojsa94/erigon/common/math"
)

var _ = (*genesisSpecMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (g Genesis) MarshalJSON() ([]byte, error) {
	type Genesis struct {
		Config                *chain.Config                                `json:"config"`
		Nonce                 math.HexOrDecimal64                          `json:"nonce"`
		Timestamp             math.HexOrDecimal64                          `json:"timestamp"`
		ExtraData             hexutility.Bytes                             `json:"extraData"`
		GasLimit              math.HexOrDecimal64                          `json:"gasLimit"   gencodec:"required"`
		Difficulty            *math.HexOrDecimal256                        `json:"difficulty" gencodec:"required"`
		Mixhash               common.Hash                                  `json:"mixHash"`
		Coinbase              common.Address                               `json:"coinbase"`
		BaseFee               *math.HexOrDecimal256                        `json:"baseFeePerGas"`
		BlobGasUsed           *math.HexOrDecimal64                         `json:"blobGasUsed"`
		ExcessBlobGas         *math.HexOrDecimal64                         `json:"excessBlobGas"`
		Alloc                 map[common0.UnprefixedAddress]GenesisAccount `json:"alloc"      gencodec:"required"`
		AuRaStep              uint64                                       `json:"auRaStep"`
		AuRaSeal              []byte                                       `json:"auRaSeal"`
		ParentBeaconBlockRoot *common.Hash                                 `json:"parentBeaconBlockRoot"`
		Number                math.HexOrDecimal64                          `json:"number"`
		GasUsed               math.HexOrDecimal64                          `json:"gasUsed"`
		ParentHash            common.Hash                                  `json:"parentHash"`
	}
	var enc Genesis
	enc.Config = g.Config
	enc.Nonce = math.HexOrDecimal64(g.Nonce)
	enc.Timestamp = math.HexOrDecimal64(g.Timestamp)
	enc.ExtraData = g.ExtraData
	enc.GasLimit = math.HexOrDecimal64(g.GasLimit)
	enc.Difficulty = (*math.HexOrDecimal256)(g.Difficulty)
	enc.Mixhash = g.Mixhash
	enc.Coinbase = g.Coinbase
	enc.BaseFee = (*math.HexOrDecimal256)(g.BaseFee)
	enc.BlobGasUsed = (*math.HexOrDecimal64)(g.BlobGasUsed)
	enc.ExcessBlobGas = (*math.HexOrDecimal64)(g.ExcessBlobGas)
	if g.Alloc != nil {
		enc.Alloc = make(map[common0.UnprefixedAddress]GenesisAccount, len(g.Alloc))
		for k, v := range g.Alloc {
			enc.Alloc[common0.UnprefixedAddress(k)] = v
		}
	}
	enc.AuRaStep = g.AuRaStep
	enc.AuRaSeal = g.AuRaSeal
	enc.ParentBeaconBlockRoot = g.ParentBeaconBlockRoot
	enc.Number = math.HexOrDecimal64(g.Number)
	enc.GasUsed = math.HexOrDecimal64(g.GasUsed)
	enc.ParentHash = g.ParentHash
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (g *Genesis) UnmarshalJSON(input []byte) error {
	type Genesis struct {
		Config                *chain.Config                                `json:"config"`
		Nonce                 *math.HexOrDecimal64                         `json:"nonce"`
		Timestamp             *math.HexOrDecimal64                         `json:"timestamp"`
		ExtraData             *hexutility.Bytes                            `json:"extraData"`
		GasLimit              *math.HexOrDecimal64                         `json:"gasLimit"   gencodec:"required"`
		Difficulty            *math.HexOrDecimal256                        `json:"difficulty" gencodec:"required"`
		Mixhash               *common.Hash                                 `json:"mixHash"`
		Coinbase              *common.Address                              `json:"coinbase"`
		BaseFee               *math.HexOrDecimal256                        `json:"baseFeePerGas"`
		BlobGasUsed           *math.HexOrDecimal64                         `json:"blobGasUsed"`
		ExcessBlobGas         *math.HexOrDecimal64                         `json:"excessBlobGas"`
		Alloc                 map[common0.UnprefixedAddress]GenesisAccount `json:"alloc"      gencodec:"required"`
		AuRaStep              *uint64                                      `json:"auRaStep"`
		AuRaSeal              []byte                                       `json:"auRaSeal"`
		ParentBeaconBlockRoot *common.Hash                                 `json:"parentBeaconBlockRoot"`
		Number                *math.HexOrDecimal64                         `json:"number"`
		GasUsed               *math.HexOrDecimal64                         `json:"gasUsed"`
		ParentHash            *common.Hash                                 `json:"parentHash"`
	}
	var dec Genesis
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Config != nil {
		g.Config = dec.Config
	}
	if dec.Nonce != nil {
		g.Nonce = uint64(*dec.Nonce)
	}
	if dec.Timestamp != nil {
		g.Timestamp = uint64(*dec.Timestamp)
	}
	if dec.ExtraData != nil {
		g.ExtraData = *dec.ExtraData
	}
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Genesis")
	}
	g.GasLimit = uint64(*dec.GasLimit)
	if dec.Difficulty == nil {
		return errors.New("missing required field 'difficulty' for Genesis")
	}
	g.Difficulty = (*big.Int)(dec.Difficulty)
	if dec.Mixhash != nil {
		g.Mixhash = *dec.Mixhash
	}
	if dec.Coinbase != nil {
		g.Coinbase = *dec.Coinbase
	}
	if dec.BaseFee != nil {
		g.BaseFee = (*big.Int)(dec.BaseFee)
	}
	if dec.BlobGasUsed != nil {
		g.BlobGasUsed = (*uint64)(dec.BlobGasUsed)
	}
	if dec.ExcessBlobGas != nil {
		g.ExcessBlobGas = (*uint64)(dec.ExcessBlobGas)
	}
	if dec.Alloc == nil {
		return errors.New("missing required field 'alloc' for Genesis")
	}
	g.Alloc = make(GenesisAlloc, len(dec.Alloc))
	for k, v := range dec.Alloc {
		g.Alloc[common.Address(k)] = v
	}
	if dec.AuRaStep != nil {
		g.AuRaStep = *dec.AuRaStep
	}
	if dec.AuRaSeal != nil {
		g.AuRaSeal = dec.AuRaSeal
	}
	if dec.ParentBeaconBlockRoot != nil {
		g.ParentBeaconBlockRoot = dec.ParentBeaconBlockRoot
	}
	if dec.Number != nil {
		g.Number = uint64(*dec.Number)
	}
	if dec.GasUsed != nil {
		g.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.ParentHash != nil {
		g.ParentHash = *dec.ParentHash
	}
	return nil
}
