package utils

import (
	"runtime"
	"sync"
)

func ProcessPairs[T any, U any](arr []T, processor func(a, b T) U) (results []U) {
	pairCountLimit := len(arr) * len(arr)
	resultsChannel := make(chan U, pairCountLimit)
	pairs := make(chan [2]T, pairCountLimit)
	// processor will be called O(N^2) times
	// so running a goroutine for each pair is not feasible
	// we run K goroutines that process each pair as they become available
	var wg sync.WaitGroup
	for i := 1; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			for pair := range pairs {
				resultsChannel <- processor(pair[0], pair[1])
			}
			wg.Done()
		}()
	}
	// load up the pairs
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			pairs <- [2]T{arr[i], arr[j]}
		}
	}
	close(pairs)
	wg.Wait()
	close(resultsChannel)
	// read results from channel and store in array
	for x := range resultsChannel {
		results = append(results, x)
	}
	return results
}
