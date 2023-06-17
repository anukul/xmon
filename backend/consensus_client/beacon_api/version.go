package beacon_api

import (
	"context"

	client "github.com/attestantio/go-eth2-client"
)

func (c *Client) Version() (string, error) {
	provider, isProvider := c.client.(client.NodeVersionProvider)
	if !isProvider {
		return "", nil
	}
	return provider.NodeVersion(context.Background())
}
