package batch

import (
	"sync"
)

func ParallelProcess[T any, Z any](items []T, concurrency int, cb func(item T) Z) []Z {
	if len(items) == 0 {
		return []Z{}
	}
	if len(items) == 1 {
		return []Z{cb(items[0])}
	}
	results := make([]Z, len(items))
	wg := sync.WaitGroup{}
	sem := make(chan struct{}, concurrency)
	for i, item := range items {
		wg.Add(1)
		sem <- struct{}{}
		go func(item T) {
			defer func() {
				wg.Done()
				<-sem
			}()
			results[i] = cb(item)
		}(item)
	}
	wg.Wait()
	return results
}
