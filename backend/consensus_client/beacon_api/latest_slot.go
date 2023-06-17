package beacon_api

import "math/big"

func (c *Client) LatestSlot() (uint64, string, error) {
	return c.Slot(big.NewInt(-1))
}
