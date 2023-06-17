package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessPairs(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	sums := ProcessPairs(a, func(a, b int) int {
		return a + b
	})
	sort.Ints(sums)
	expectedSums := []int{3, 4, 5, 5, 6, 6, 7, 7, 8, 9}
	assert.Equal(t, expectedSums, sums)
}
