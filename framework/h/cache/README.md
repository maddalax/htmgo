# Pluggable Cache System for htmgo

## Overview

The htmgo framework now supports a pluggable cache system that allows developers to provide their own caching
implementations. This addresses potential memory exhaustion vulnerabilities in the previous TTL-only caching approach
and provides greater flexibility for production deployments.

## Motivation

The previous caching mechanism relied exclusively on Time-To-Live (TTL) expiration, which could lead to:

- **Unbounded memory growth**: High-cardinality cache keys could consume all available memory
- **DDoS vulnerability**: Attackers could exploit this by generating many unique cache keys
- **Limited flexibility**: No support for size-bounded caches or distributed caching solutions

## Architecture

The new system introduces a generic `Store[K comparable, V any]` interface:

```go
package main

import "time"

type Store[K comparable, V any] interface {
	// Set adds or updates an entry in the cache with the given TTL
	Set(key K, value V, ttl time.Duration)

	// GetOrCompute atomically gets an existing value or computes and stores a new value
	// This prevents duplicate computation when multiple goroutines request the same key
	GetOrCompute(key K, compute func() V, ttl time.Duration) V

	// Delete removes an entry from the cache
	Delete(key K)

	// Purge removes all items from the cache
	Purge()

	// Close releases any resources used by the cache
	Close()
}
```

### Atomic Guarantees

The `GetOrCompute` method provides **atomic guarantees** to prevent cache stampedes and duplicate computations:
- When multiple goroutines request the same uncached key simultaneously, only one will execute the compute function
- Other goroutines will wait and receive the computed result
- This eliminates race conditions that could cause duplicate expensive operations like database queries or renders

## Usage

### Using the Default Cache

By default, htmgo continues to use a TTL-based cache for backward compatibility:

```go
// No changes needed - works exactly as before
UserProfile := h.CachedPerKeyT(
  15*time.Minute,
  func(userID int) (int, h.GetElementFunc) {
	return userID, func() *h.Element {
		return h.Div(h.Text("User profile"))
	}
  },
)
```

### Using a Custom Cache

You can provide your own cache implementation using the `WithStore` option:

```go
package main

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/h/cache"
	"time"
)

var (
	// Create a memory-bounded LRU cache
	lruCache = cache.NewLRUStore[any, string](10_000) // Max 10,000 items

	// Use it with a cached component
	UserProfile = h.CachedPerKeyT(
		15*time.Minute,
		func (userID int) (int, h.GetElementFunc) {
			return userID, func () *h.Element {
				return h.Div(h.Text("User profile"))
			}
		},
		h.WithStore(lruCache), // Pass the custom cache
	)
)
```

### Changing the Default Cache Globally

You can override the default cache provider for your entire application:

```go
package main

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/h/cache"
)

func init() {
	// All cached components will use LRU by default
	h.DefaultCacheProvider = func () cache.Store[any, string] {
		return cache.NewLRUStore[any, string](50_000)
	}
}
```

## Example Implementations

### Built-in Stores

1. **TTLStore** (default): Time-based expiration with periodic cleanup
2. **LRUStore** (example): Least Recently Used eviction with size limits

### Integrating Third-Party Libraries

Here's an example of integrating the high-performance `go-freelru` library:

```go
import (
  "time"
  "github.com/elastic/go-freelru"
  "github.com/maddalax/htmgo/framework/h/cache"
)

type FreeLRUAdapter[K comparable, V any] struct {
lru *freelru.LRU[K, V]
}

func NewFreeLRUAdapter[K comparable, V any](size uint32) cache.Store[K, V] {
lru, err := freelru.New[K, V](size, nil)
if err != nil {
panic(err)
}
return &FreeLRUAdapter[K, V]{lru: lru}
}

func (s *FreeLRUAdapter[K, V]) Set(key K, value V, ttl time.Duration) {
// Note: go-freelru doesn't support per-item TTL
s.lru.Add(key, value)
}

func (s *FreeLRUAdapter[K, V]) GetOrCompute(key K, compute func() V, ttl time.Duration) V {
    // Check if exists in cache
    if val, ok := s.lru.Get(key); ok {
        return val
    }
    
    // Not in cache, compute and store
    // Note: This simple implementation doesn't provide true atomic guarantees
    // For production use, you'd need additional synchronization
    value := compute()
    s.lru.Add(key, value)
    return value
}

func (s *FreeLRUAdapter[K, V]) Delete(key K) {
s.lru.Remove(key)
}

func (s *FreeLRUAdapter[K, V]) Purge() {
s.lru.Clear()
}

func (s *FreeLRUAdapter[K, V]) Close() {
// No-op for this implementation
}
```

### Redis-based Distributed Cache

```go
type RedisStore struct {
client *redis.Client
prefix string
}

func (s *RedisStore) Set(key any, value string, ttl time.Duration) {
keyStr := fmt.Sprintf("%s:%v", s.prefix, key)
s.client.Set(context.Background(), keyStr, value, ttl)
}

func (s *RedisStore) GetOrCompute(key any, compute func() string, ttl time.Duration) string {
    keyStr := fmt.Sprintf("%s:%v", s.prefix, key)
    ctx := context.Background()
    
    // Try to get from Redis
    val, err := s.client.Get(ctx, keyStr).Result()
    if err == nil {
        return val
    }
    
    // Not in cache, compute new value
    // For true atomic guarantees, use Redis SET with NX option
    value := compute()
    s.client.Set(ctx, keyStr, value, ttl)
    return value
}

// ... implement other methods
```

## Migration Guide

### For Existing Applications

The changes are backward compatible. Existing applications will continue to work without modifications. The function
signatures now accept optional `CacheOption` parameters, but these can be omitted.

### Recommended Migration Path

1. **Assess your caching needs**: Determine if you need memory bounds or distributed caching
2. **Choose an implementation**: Use the built-in LRUStore or integrate a third-party library
3. **Update critical components**: Start with high-traffic or high-cardinality cached components
4. **Monitor memory usage**: Ensure your cache size limits are appropriate

## Security Considerations

### Memory-Bounded Caches

For public-facing applications, we strongly recommend using a memory-bounded cache to prevent DoS attacks:

```go
// Limit cache to reasonable size based on your server's memory
cache := cache.NewLRUStore[any, string](100_000)

// Use for all user-specific caching
UserContent := h.CachedPerKey(
5*time.Minute,
getUserContent,
h.WithStore(cache),
)
```

### Cache Key Validation

When using user input as cache keys, always validate and sanitize:

```go
func cacheKeyForUser(userInput string) string {
// Limit length and remove special characters
key := strings.TrimSpace(userInput)
if len(key) > 100 {
key = key[:100]
}
return regexp.MustCompile(`[^a-zA-Z0-9_-]`).ReplaceAllString(key, "")
}
```

## Performance Considerations

1. **TTLStore**: Best for small caches with predictable key patterns
2. **LRUStore**: Good general-purpose choice with memory bounds  
3. **Third-party stores**: Consider `go-freelru` or `theine-go` for high-performance needs
4. **Distributed stores**: Use Redis/Memcached for multi-instance deployments
5. **Atomic Operations**: The `GetOrCompute` method prevents duplicate computations, significantly improving performance under high concurrency

### Concurrency Benefits

The atomic `GetOrCompute` method provides significant performance benefits:
- **Prevents Cache Stampedes**: When a popular cache entry expires, only one goroutine will recompute it
- **Reduces Load**: Expensive operations (database queries, API calls, complex renders) are never duplicated
- **Improves Response Times**: Waiting goroutines get results faster than computing themselves

## Best Practices

1. **Set appropriate cache sizes**: Balance memory usage with hit rates
2. **Use consistent TTLs**: Align with your data update patterns
3. **Monitor cache metrics**: Track hit rates, evictions, and memory usage
4. **Handle cache failures gracefully**: Caches should enhance, not break functionality
5. **Close caches properly**: Call `Close()` during graceful shutdown
6. **Implement atomic guarantees**: Ensure your `GetOrCompute` implementation prevents concurrent computation
7. **Test concurrent access**: Verify your cache handles simultaneous requests correctly

## Future Enhancements

- Built-in metrics and monitoring hooks
- Automatic size estimation for cached values
- Warming and preloading strategies
- Cache invalidation patterns