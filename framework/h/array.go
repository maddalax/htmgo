package h

import (
	"github.com/maddalax/htmgo/framework/datastructure/orderedmap"
)

// Unique returns a new slice with only unique items.
func Unique[T any](slice []T, key func(item T) string) []T {
	var result []T
	seen := make(map[string]bool)
	for _, v := range slice {
		k := key(v)
		if _, ok := seen[k]; !ok {
			seen[k] = true
			result = append(result, v)
		}
	}
	return result
}

// Find returns the first item in the slice that matches the predicate.
func Find[T any](slice []T, predicate func(item *T) bool) *T {
	for _, v := range slice {
		if predicate(&v) {
			return &v
		}
	}
	return nil
}

// GroupBy groups the items in the slice by the key returned by the key function.
func GroupBy[T any, K comparable](slice []T, key func(item T) K) map[K][]T {
	grouped := make(map[K][]T)
	for _, item := range slice {
		k := key(item)
		items, ok := grouped[k]
		if !ok {
			items = []T{}
		}
		grouped[k] = append(items, item)
	}
	return grouped
}

// GroupByOrdered groups the items in the slice by the key returned by the key function, and returns an Map.
func GroupByOrdered[T any, K comparable](slice []T, key func(item T) K) *orderedmap.Map[K, []T] {
	grouped := orderedmap.New[K, []T]()
	for _, item := range slice {
		k := key(item)
		items, ok := grouped.Get(k)
		if !ok {
			items = []T{}
		}
		grouped.Set(k, append(items, item))
	}
	return grouped
}

// Filter returns a new slice with only items that match the predicate.
func Filter[T any](slice []T, predicate func(item T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map returns a new slice with the results of the mapper function.
func Map[T, U any](slice []T, mapper func(item T) U) []U {
	var result []U
	for _, v := range slice {
		result = append(result, mapper(v))
	}
	return result
}
