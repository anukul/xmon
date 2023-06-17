package beacon_api

import (
	"context"
	"github.com/rs/zerolog"
	"time"

	eth2client "github.com/attestantio/go-eth2-client"
	ehttp "github.com/attestantio/go-eth2-client/http"
)

type Client struct {
	client eth2client.Service
}

func NewClient(url string, headers map[string]string) (*Client, error) {
	client, err := ehttp.New(context.Background(),
		ehttp.WithAddress(url),
		ehttp.WithExtraHeaders(headers),
		ehttp.WithLogLevel(zerolog.InfoLevel),
		ehttp.WithTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	return &Client{client: client}, nil
}
