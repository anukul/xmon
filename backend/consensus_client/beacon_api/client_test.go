package beacon_api

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

var client *Client

func setup() func() {
	var err error
	client, err = NewClient(
		"https://prysm.mainnet.ethpandaops.io",
		map[string]string{
			"CF-Access-Client-Id":     "",
			"CF-Access-Client-Secret": "",
		})
	if err != nil {
		panic(err)
	}
	return func() {}
}

func TestBeaconNode_GetSlot(t *testing.T) {
	defer setup()()
	slotNumber := 6683720
	slot, root, err := client.Slot(big.NewInt(int64(slotNumber)))
	if err != nil {
		panic(err)
	}
	assert.Equal(t, uint64(slotNumber), slot)
	assert.Equal(t, "0x18613a4456542c2f73f8507297939bb18c43bc5fa3f810b671a31376e673ad22", root)
}

func TestBeaconNode_GetLatestSlot(t *testing.T) {
	defer setup()()
	slot, root, err := client.LatestSlot()
	if err != nil {
		panic(err)
	}
	fmt.Println(slot, root)
}
