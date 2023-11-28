package spectest

import (
	"os"
	"testing"

	"github.com/nebojsa94/erigon/spectest"

	"github.com/nebojsa94/erigon/cl/transition"

	"github.com/nebojsa94/erigon/cl/spectest/consensus_tests"
)

func Test(t *testing.T) {
	spectest.RunCases(t, consensus_tests.TestFormats, transition.ValidatingMachine, os.DirFS("./tests"))
}
