package json_rpc

import "context"

func (c *Client) Version() (string, error) {
	var version string
	err := c.rpc.CallContext(context.Background(), &version, "web3_clientVersion")
	return version, err
}
