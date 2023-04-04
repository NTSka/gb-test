package calc

import (
	"context"
	"emperror.dev/errors"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
	"log"
)

type Service struct {
	client IClient
	cfg    *Config
}

func NewService(cfg *Config, client IClient) IService {
	return &Service{client: client, cfg: cfg}
}

func (s *Service) GetMaxDiff(ctx context.Context) (string, error) {
	last, err := s.client.GetLastBlock(ctx)
	if err != nil {
		return "", errors.Wrap(err, "calcClient.GetLastBlock")
	}

	log.Printf("Last block: %d\n", last)

	diffs := newDiffMap()

	errGroup, ctx := errgroup.WithContext(ctx)

	errGroup.SetLimit(s.cfg.Parallel)

	for i := last; i > last-s.cfg.Count; i-- {
		i := i
		errGroup.Go(func() error {
			err := s.getBlockDiffs(ctx, diffs, i)
			if err != nil {
				return err
			}

			log.Printf("Done %d\n", i)

			return nil
		})
	}

	if err = errGroup.Wait(); err != nil {
		return "", errors.Wrap(err, "calcSvc.getBlockDiffs")
	}

	maxValue := decimal.NewFromInt(0)
	maxAddress := ""

	for address, diff := range diffs.diffs {
		absDiff := diff.Abs()
		if absDiff.GreaterThan(maxValue) {
			maxValue = absDiff
			maxAddress = address
		}
	}

	return maxAddress, nil
}

func (s *Service) getBlockDiffs(ctx context.Context, diffs *diffMap, blockNumber int64) error {
	block, err := s.client.GetBlockByNumber(ctx, blockNumber)
	if err != nil {
		return errors.Wrap(err, "calcClient.GetBlockByNumber")
	}

	for _, tx := range block.Transactions() {
		signer := types.NewLondonSigner(tx.ChainId())
		from, err := types.Sender(signer, tx)
		if err != nil {
			return errors.Wrap(err, "ethTypes.Sender")
		}

		to := tx.To()

		amountTo := decimal.NewFromBigInt(tx.Value(), 0)
		gas := decimal.NewFromBigInt(tx.GasPrice(), 0).Mul(decimal.NewFromInt(int64(tx.Gas())))
		amountFrom := amountTo.Add(gas)

		diffs.addValue(from.Hex(), amountFrom)
		if to != nil {
			diffs.addValue(to.Hex(), amountTo)
		}
	}

	return nil
}
