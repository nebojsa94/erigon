package snapshot_format

import (
	"io"

	libcommon "github.com/nebojsa94/erigon/erigon-lib/common"
	"github.com/nebojsa94/erigon/cl/cltypes"
)

type MockBlockReader struct {
	Block *cltypes.Eth1Block
}

func (t *MockBlockReader) WithdrawalsSZZ(out io.Writer, number uint64, hash libcommon.Hash) error {
	l, err := t.Block.Withdrawals.EncodeSSZ(nil)
	if err != nil {
		return err
	}
	_, err = out.Write(l)
	return err
}

func (t *MockBlockReader) TransactionsSSZ(out io.Writer, number uint64, hash libcommon.Hash) error {
	l, err := t.Block.Transactions.EncodeSSZ(nil)
	if err != nil {
		return err
	}
	_, err = out.Write(l)
	return err
}
