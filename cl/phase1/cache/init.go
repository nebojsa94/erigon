package cache

import (
	"github.com/nebojsa94/erigon-lib/common"
	"github.com/nebojsa94/erigon/cl/phase1/core/state/lru"
)

func init() {
	var err error
	if attestationIndiciesCache, err = lru.New[common.Hash, []uint64]("attestationIndiciesCacheSize", attestationIndiciesCacheSize); err != nil {
		panic(err)
	}
}
