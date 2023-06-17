package utils

import (
	"sort"
)

func BinarySearch(low, high int, fn func(current int) bool) int {
	ans := sort.Search(high-low, func(i int) bool {
		return fn(i + low)
	})
	return ans + low
}
