package cache

import (
	"time"
)

// Store defines the interface for a pluggable cache.
// This allows users to provide their own caching implementations, such as LRU, LFU,
// or even distributed caches. The cache implementation is responsible for handling
// its own eviction policies (TTL, size limits, etc.).
type Store[K comparable, V any] interface {
	// Set adds or updates an entry in the cache. The implementation should handle the TTL.
	Set(key K, value V, ttl time.Duration)

	// GetOrCompute atomically gets an existing value or computes and stores a new value.
	// This method prevents duplicate computation when multiple goroutines request the same key.
	// The compute function is called only if the key is not found or has expired.
	GetOrCompute(key K, compute func() V, ttl time.Duration) V

	// Delete removes an entry from the cache.
	Delete(key K)

	// Purge removes all items from the cache.
	Purge()

	// Close releases any resources used by the cache, such as background goroutines.
	Close()
}
