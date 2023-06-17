package utils

import (
	"fmt"

	"github.com/anukul/xmon/backend/constants"
)

func BlockHashKey(blockNumber int64) string {
	return fmt.Sprintf("%s_%d", constants.BlockHash, blockNumber)
}

func BlockNumberKey() string {
	return constants.BlockNumber
}
