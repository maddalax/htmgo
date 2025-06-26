package cache

import (
	"flag"
	"log/slog"
	"sync"
	"time"
)

// TTLStore is a time-to-live based cache implementation that mimics
// the original htmgo caching behavior. It stores values with expiration
// times and periodically cleans up expired entries.
type TTLStore[K comparable, V any] struct {
	cache     map[K]*entry[V]
	mutex     sync.RWMutex
	closeOnce sync.Once
	closeChan chan struct{}
}

type entry[V any] struct {
	value      V
	expiration time.Time
}

// NewTTLStore creates a new TTL-based cache store.
func NewTTLStore[K comparable, V any]() Store[K, V] {
	s := &TTLStore[K, V]{
		cache:     make(map[K]*entry[V]),
		closeChan: make(chan struct{}),
	}
	s.startCleaner()
	return s
}

// Set adds or updates an entry in the cache with the given TTL.
func (s *TTLStore[K, V]) Set(key K, value V, ttl time.Duration) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache[key] = &entry[V]{
		value:      value,
		expiration: time.Now().Add(ttl),
	}
}

// Get retrieves an entry from the cache.
func (s *TTLStore[K, V]) Get(key K) (V, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var zero V
	e, ok := s.cache[key]
	if !ok {
		return zero, false
	}

	// Check if expired
	if time.Now().After(e.expiration) {
		return zero, false
	}

	return e.value, true
}

// Delete removes an entry from the cache.
func (s *TTLStore[K, V]) Delete(key K) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.cache, key)
}

// Purge removes all items from the cache.
func (s *TTLStore[K, V]) Purge() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache = make(map[K]*entry[V])
}

// Close stops the background cleaner goroutine.
func (s *TTLStore[K, V]) Close() {
	s.closeOnce.Do(func() {
		close(s.closeChan)
	})
}

// startCleaner starts a background goroutine that periodically removes expired entries.
func (s *TTLStore[K, V]) startCleaner() {
	isTests := flag.Lookup("test.v") != nil

	go func() {
		ticker := time.NewTicker(time.Minute)
		if isTests {
			ticker = time.NewTicker(time.Second)
		}
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.clearExpired()
			case <-s.closeChan:
				return
			}
		}
	}()
}

// clearExpired removes all expired entries from the cache.
func (s *TTLStore[K, V]) clearExpired() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now()
	deletedCount := 0

	for key, e := range s.cache {
		if now.After(e.expiration) {
			delete(s.cache, key)
			deletedCount++
		}
	}

	if deletedCount > 0 {
		slog.Debug("Deleted expired cache entries", slog.Int("count", deletedCount))
	}
}
