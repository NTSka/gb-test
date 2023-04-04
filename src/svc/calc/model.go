package calc

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"sync"
)

type IService interface {
	GetMaxDiff(ctx context.Context) (string, error)
}

type IClient interface {
	GetLastBlock(ctx context.Context) (int64, error)
	GetBlockByNumber(ctx context.Context, blockNumber int64) (*types.Block, error)
}

type diffMap struct {
	sync.Mutex
	diffs map[string]decimal.Decimal
}

func (d *diffMap) addValue(address string, amount decimal.Decimal) {
	d.Lock()
	defer d.Unlock()

	d.diffs[address] = d.diffs[address].Add(amount)
}

func newDiffMap() *diffMap {
	return &diffMap{
		sync.Mutex{},
		make(map[string]decimal.Decimal),
	}
}
