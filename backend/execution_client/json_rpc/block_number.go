package json_rpc

import (
	"context"
	"math/big"

	"github.com/anukul/xmon/backend/utils"
)

func (c *Client) BlockNumber() (*big.Int, error) {
	var blockNumber big.Int
	if err := c.db.Get(utils.BlockNumberKey(), &blockNumber); err != nil {
		return nil, err
	}
	return &blockNumber, nil
}

func (c *Client) GetBlockNumber() (*big.Int, error) {
	_blockNumber, err := c.eth.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	blockNumber := big.NewInt(int64(_blockNumber))
	if err := c.db.Set(utils.BlockNumberKey(), blockNumber); err != nil {
		return nil, err
	}
	return blockNumber, nil
}
