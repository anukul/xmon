package el_test

import (
	"github.com/anukul/xmon/backend/common"
	"github.com/anukul/xmon/backend/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/syndtr/goleveldb/leveldb"
	"math/big"
)

type Chain map[int64]string

type Client struct {
	chain Chain
}

func NewClient(c Chain, db common.DatabaseClient) *Client {
	for blockNumber, blockHash := range c {
		if err := db.Set(utils.BlockHashKey(blockNumber), blockHash); err != nil {
			panic(err)
		}
	}

	return &Client{
		chain: c,
	}
}

func (c *Client) BlockHash(blockNumber *big.Int) (string, error) {
	res, ok := c.chain[blockNumber.Int64()]
	if !ok {
		return "", leveldb.ErrNotFound
	}
	return res, nil
}

func (c *Client) GetBlockHash(blockNumber *big.Int) (string, error) {
	blockHash, err := c.BlockHash(blockNumber)
	// function is expected to fetch from client if hash not found in db
	if err == leveldb.ErrNotFound {
		err = ethereum.NotFound
	}
	return blockHash, err
}

func (c *Client) BlockNumber() (*big.Int, error) {
	head := int64(0)
	for block, _ := range c.chain {
		if block > head {
			head = block
		}
	}
	return big.NewInt(head), nil
}

func (c *Client) GetBlockNumber() (*big.Int, error) {
	return c.BlockNumber()
}

func (c *Client) Version() (string, error) {
	return "test", nil
}
