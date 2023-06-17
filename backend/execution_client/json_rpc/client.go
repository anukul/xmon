package json_rpc

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/anukul/xmon/backend/common"
)

type Client struct {
	logger *zerolog.Logger
	rpc    *rpc.Client
	eth    *ethclient.Client
	db     common.DatabaseClient
}

func NewClient(url string, headers map[string]string, db common.DatabaseClient) (*Client, error) {
	logger := log.With().Str("component", "execution_client").Logger()
	_headers := make(http.Header)
	for k, v := range headers {
		_headers.Set(k, v)
	}
	rpcClient, err := rpc.DialOptions(context.Background(), url, rpc.WithHeaders(_headers))
	if err != nil {
		return nil, err
	}
	ethClient := ethclient.NewClient(rpcClient)
	return &Client{
		rpc:    rpcClient,
		eth:    ethClient,
		db:     db,
		logger: &logger,
	}, nil
}
