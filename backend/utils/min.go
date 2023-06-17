package utils

import (
	"math/big"
)

func MinBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) < 0 {
		return a
	} else {
		return b
	}
}

func MaxBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) > 0 {
		return a
	} else {
		return b
	}
}
