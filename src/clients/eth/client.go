package eth

import (
	"context"
	"emperror.dev/errors"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Client struct {
	c *ethclient.Client
}

func NewClient(c *Config) (*Client, error) {
	ethClient, err := ethclient.Dial(c.URI)
	if err != nil {
		return nil, errors.Wrap(err, "ethclient.Dial")
	}
	return &Client{
		ethClient,
	}, nil
}

func (c *Client) GetLastBlock(ctx context.Context) (int64, error) {
	u, err := c.c.BlockNumber(ctx)
	return int64(u), err
}

func (c *Client) GetBlockByNumber(ctx context.Context, blockNumber int64) (*types.Block, error) {
	return c.c.BlockByNumber(ctx, big.NewInt(blockNumber))
}
