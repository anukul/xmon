package beacon_api

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	eth2client "github.com/attestantio/go-eth2-client"
)

func (c *Client) Slot(slotNumber *big.Int) (uint64, string, error) {
	var blockID string
	if slotNumber.Cmp(big.NewInt(-1)) == 0 {
		blockID = "head"
	} else {
		blockID = slotNumber.String()
	}

	provider, isProvider := c.client.(eth2client.BeaconBlockHeadersProvider)
	if !isProvider {
		return 0, "", errors.New("err")
	}
	beaconBlockHeader, err := provider.BeaconBlockHeader(context.Background(), blockID)
	if err != nil {
		return 0, "", err
	}
	slot := beaconBlockHeader.Header.Message.Slot
	slotRoot := fmt.Sprintf("0x%s", hex.EncodeToString(beaconBlockHeader.Root[:]))
	return uint64(slot), slotRoot, nil
}
