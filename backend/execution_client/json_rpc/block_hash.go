package json_rpc

import (
	"context"
	"math/big"

	"github.com/anukul/xmon/backend/utils"
)

func (c *Client) BlockHash(blockNumber *big.Int) (string, error) {
	var blockHash string
	// fetch from db
	if err := c.db.Get(utils.BlockHashKey(blockNumber.Int64()), &blockHash); err != nil {
		return "", err
	}
	return blockHash, nil
}

func (c *Client) GetBlockHash(blockNumber *big.Int) (string, error) {
	// check db
	if blockHash, err := c.BlockHash(blockNumber); err == nil {
		return blockHash, nil
	}
	// fetch from node
	blockHeader, err := c.eth.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		return "", err
	}
	blockHash := blockHeader.Hash().String()
	// save to db
	if err = c.db.Set(utils.BlockHashKey(blockNumber.Int64()), blockHash); err != nil {
		return "", err
	}
	return blockHash, nil
}
