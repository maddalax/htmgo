package cache

import (
	"container/list"
	"sync"
	"time"
)

// LRUStore is an example of a memory-bounded cache implementation using
// the Least Recently Used (LRU) eviction policy. This demonstrates how
// to create a custom cache store that prevents unbounded memory growth.
//
// This is a simple example implementation. For production use, consider
// using optimized libraries like github.com/elastic/go-freelru or
// github.com/Yiling-J/theine-go.
type LRUStore[K comparable, V any] struct {
	maxSize   int
	cache     map[K]*list.Element
	lru       *list.List
	mutex     sync.RWMutex
	closeChan chan struct{}
	closeOnce sync.Once
}

type lruEntry[K comparable, V any] struct {
	key        K
	value      V
	expiration time.Time
}

// NewLRUStore creates a new LRU cache with the specified maximum size.
// When the cache reaches maxSize, the least recently used items are evicted.
func NewLRUStore[K comparable, V any](maxSize int) Store[K, V] {
	if maxSize <= 0 {
		panic("LRUStore maxSize must be positive")
	}

	s := &LRUStore[K, V]{
		maxSize:   maxSize,
		cache:     make(map[K]*list.Element),
		lru:       list.New(),
		closeChan: make(chan struct{}),
	}

	// Start a goroutine to periodically clean up expired entries
	go s.cleanupExpired()

	return s
}

// Set adds or updates an entry in the cache with the given TTL.
// If the cache is at capacity, the least recently used item is evicted.
func (s *LRUStore[K, V]) Set(key K, value V, ttl time.Duration) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	expiration := time.Now().Add(ttl)

	// Check if key already exists
	if elem, exists := s.cache[key]; exists {
		// Update existing entry and move to front
		entry := elem.Value.(*lruEntry[K, V])
		entry.value = value
		entry.expiration = expiration
		s.lru.MoveToFront(elem)
		return
	}

	// Add new entry
	entry := &lruEntry[K, V]{
		key:        key,
		value:      value,
		expiration: expiration,
	}
	elem := s.lru.PushFront(entry)
	s.cache[key] = elem

	// Evict oldest if over capacity
	if s.lru.Len() > s.maxSize {
		oldest := s.lru.Back()
		if oldest != nil {
			s.removeElement(oldest)
		}
	}
}

// Get retrieves an entry from the cache.
// Returns the value and true if found and not expired, zero value and false otherwise.
func (s *LRUStore[K, V]) Get(key K) (V, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var zero V
	elem, exists := s.cache[key]
	if !exists {
		return zero, false
	}

	entry := elem.Value.(*lruEntry[K, V])

	// Check if expired
	if time.Now().After(entry.expiration) {
		s.removeElement(elem)
		return zero, false
	}

	// Move to front (mark as recently used)
	s.lru.MoveToFront(elem)
	return entry.value, true
}

// Delete removes an entry from the cache.
func (s *LRUStore[K, V]) Delete(key K) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if elem, exists := s.cache[key]; exists {
		s.removeElement(elem)
	}
}

// Purge removes all items from the cache.
func (s *LRUStore[K, V]) Purge() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache = make(map[K]*list.Element)
	s.lru.Init()
}

// Close stops the background cleanup goroutine.
func (s *LRUStore[K, V]) Close() {
	s.closeOnce.Do(func() {
		close(s.closeChan)
	})
}

// removeElement removes an element from both the map and the list.
// Must be called with the mutex held.
func (s *LRUStore[K, V]) removeElement(elem *list.Element) {
	entry := elem.Value.(*lruEntry[K, V])
	delete(s.cache, entry.key)
	s.lru.Remove(elem)
}

// cleanupExpired periodically removes expired entries.
func (s *LRUStore[K, V]) cleanupExpired() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.removeExpired()
		case <-s.closeChan:
			return
		}
	}
}

// removeExpired scans the cache and removes expired entries.
func (s *LRUStore[K, V]) removeExpired() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now()
	// Create a slice to hold elements to remove to avoid modifying list during iteration
	var toRemove []*list.Element

	for elem := s.lru.Back(); elem != nil; elem = elem.Prev() {
		entry := elem.Value.(*lruEntry[K, V])
		if now.After(entry.expiration) {
			toRemove = append(toRemove, elem)
		}
	}

	// Remove expired elements
	for _, elem := range toRemove {
		s.removeElement(elem)
	}
}
