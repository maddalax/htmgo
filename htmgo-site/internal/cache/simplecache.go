package cache

import (
	"sync"
	"time"
)

type SimpleCache struct {
	data map[string]Entry
	lock sync.RWMutex
}

type Entry struct {
	Value      any
	Expiration time.Time
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{
		data: make(map[string]Entry),
		lock: sync.RWMutex{},
	}
}

func (c *SimpleCache) Get(key string) (any, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	if entry.Expiration.Before(time.Now()) {
		delete(c.data, key)
		return nil, false
	}
	return entry.Value, true
}

func (c *SimpleCache) Set(key string, value any, expiration time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = Entry{
		Value:      value,
		Expiration: time.Now().Add(expiration),
	}
}

func GetOrSet[T any](cache *SimpleCache, key string, expiration time.Duration, cb func() (T, bool)) T {
	if val, ok := cache.Get(key); ok {
		return val.(T)
	}
	value, should := cb()
	if should {
		cache.Set(key, value, expiration)
	}
	return value
}
