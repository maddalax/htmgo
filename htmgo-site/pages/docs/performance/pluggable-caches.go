package performance

import (
	"github.com/maddalax/htmgo/framework/h"
	. "htmgo-site/pages/docs"
	"htmgo-site/ui"
)

func PluggableCaches(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Creating Custom Cache Stores"),
			Text(`
				htmgo now supports pluggable cache stores, allowing you to use any caching backend or implement custom caching strategies.
				This feature enables better control over memory usage, distributed caching support, and protection against memory exhaustion attacks.
			`),

			SubTitle("The Cache Store Interface"),
			Text(`
				All cache stores implement the following interface:
			`),
			ui.GoCodeSnippet(CacheStoreInterface),
			Text(`
				The interface is generic, supporting any comparable key type and any value type.
				The <b>Close()</b> method allows for cleanup of resources when the cache is no longer needed.
			`),

			SubTitle("Using Custom Cache Stores"),
			Text(`
				You can use custom cache stores in two ways:
			`),
			StepTitle("1. Per-Component Configuration"),
			ui.GoCodeSnippet(PerComponentExample),

			StepTitle("2. Global Default Configuration"),
			ui.GoCodeSnippet(GlobalConfigExample),

			SubTitle("Implementing a Custom Cache Store"),
			Text(`
				Here's a complete example of implementing a Redis-based cache store:
			`),
			ui.GoCodeSnippet(RedisCacheExample),

			SubTitle("Built-in Cache Stores"),
			Text(`
				htmgo provides two built-in cache implementations:
			`),

			StepTitle("TTL Store (Default)"),
			Text(`
				The default cache store that maintains backward compatibility with existing htmgo applications.
				It automatically removes expired entries based on TTL.
			`),
			ui.GoCodeSnippet(TTLStoreExample),

			StepTitle("LRU Store"),
			Text(`
				A memory-bounded cache that evicts least recently used items when the size limit is reached.
				This is useful for preventing memory exhaustion attacks.
			`),
			ui.GoCodeSnippet(LRUStoreExample),

			SubTitle("Migration Guide"),
			Text(`
				<b>Good news!</b> Existing htmgo applications require <b>no changes</b> to work with the new cache system.
				The default behavior remains exactly the same. However, if you want to take advantage of the new features:
			`),

			StepTitle("Before (existing code):"),
			ui.GoCodeSnippet(MigrationBefore),

			StepTitle("After (with custom cache):"),
			ui.GoCodeSnippet(MigrationAfter),

			SubTitle("Best Practices"),
			Text(`
				<b>1. Resource Management:</b> Always implement the Close() method if your cache uses external resources.
			`),
			Text(`
				<b>2. Thread Safety:</b> Ensure your cache implementation is thread-safe as it will be accessed concurrently.
			`),
			Text(`
				<b>3. Memory Bounds:</b> Consider implementing size limits to prevent unbounded memory growth.
			`),
			Text(`
				<b>4. Error Handling:</b> Cache operations should be resilient to failures and not crash the application.
			`),
			Text(`
				<b>5. Monitoring:</b> Consider adding metrics to track cache hit rates and performance.
			`),

			SubTitle("Common Use Cases"),

			StepTitle("Distributed Caching"),
			Text(`
				Use Redis or Memcached for sharing cache across multiple application instances:
			`),
			ui.GoCodeSnippet(DistributedCacheExample),

			StepTitle("Memory-Bounded Caching"),
			Text(`
				Prevent memory exhaustion by limiting cache size:
			`),
			ui.GoCodeSnippet(MemoryBoundedExample),

			StepTitle("Tiered Caching"),
			Text(`
				Implement a multi-level cache with fast local storage and slower distributed storage:
			`),
			ui.GoCodeSnippet(TieredCacheExample),

			Text(`
				<b>Security Note:</b> The pluggable cache system helps mitigate memory exhaustion attacks by allowing
				you to implement bounded caches. Always consider using size-limited caches in production environments
				where untrusted input could influence cache keys.
			`),

			NextStep(
				"mt-4",
				PrevBlock("Caching Per Key", DocPath("/performance/caching-per-key")),
				NextBlock("Server Sent Events", DocPath("/pushing-data/sse")),
			),
		),
	)
}

const CacheStoreInterface = `
type Store[K comparable, V any] interface {
    Get(key K) (V, bool)
    Set(key K, value V)
    Delete(key K)
    Close() error
}
`

const PerComponentExample = `
// Create a custom cache store
lruCache := cache.NewLRUStore[string, string](10000) // Max 10k items

// Use it with a cached component
var CachedUserProfile = h.CachedPerKeyT(
    15*time.Minute,
    getUserProfile,
    h.WithStore(lruCache), // Pass the custom store
)
`

const GlobalConfigExample = `
// Set a global default cache provider
func init() {
    h.DefaultCacheProvider = func() cache.Store[any, string] {
        return cache.NewLRUStore[any, string](50000)
    }
}

// All cached components will now use LRU caching by default
var CachedData = h.Cached(5*time.Minute, getData) // Uses LRU store
`

const RedisCacheExample = `
package cache

import (
    "context"
    "encoding/json"
    "time"
    "github.com/redis/go-redis/v9"
)

type RedisStore[K comparable, V any] struct {
    client *redis.Client
    prefix string
    ttl    time.Duration
}

func NewRedisStore[K comparable, V any](client *redis.Client, prefix string, ttl time.Duration) *RedisStore[K, V] {
    return &RedisStore[K, V]{
        client: client,
        prefix: prefix,
        ttl:    ttl,
    }
}

func (r *RedisStore[K, V]) Get(key K) (V, bool) {
    var zero V
    ctx := context.Background()
    
    // Create Redis key
    redisKey := fmt.Sprintf("%s:%v", r.prefix, key)
    
    // Get value from Redis
    data, err := r.client.Get(ctx, redisKey).Bytes()
    if err != nil {
        return zero, false
    }
    
    // Deserialize value
    var value V
    if err := json.Unmarshal(data, &value); err != nil {
        return zero, false
    }
    
    return value, true
}

func (r *RedisStore[K, V]) Set(key K, value V) {
    ctx := context.Background()
    redisKey := fmt.Sprintf("%s:%v", r.prefix, key)
    
    // Serialize value
    data, err := json.Marshal(value)
    if err != nil {
        return
    }
    
    // Set in Redis with TTL
    r.client.Set(ctx, redisKey, data, r.ttl)
}

func (r *RedisStore[K, V]) Delete(key K) {
    ctx := context.Background()
    redisKey := fmt.Sprintf("%s:%v", r.prefix, key)
    r.client.Del(ctx, redisKey)
}

func (r *RedisStore[K, V]) Close() error {
    return r.client.Close()
}

// Usage
redisClient := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})

redisCache := NewRedisStore[string, string](
    redisClient,
    "myapp:cache",
    15*time.Minute,
)

var CachedUserData = h.CachedPerKeyT(
    15*time.Minute,
    getUserData,
    h.WithStore(redisCache),
)
`

const TTLStoreExample = `
// Create a TTL-based cache (this is the default)
ttlCache := cache.NewTTLStore[string, string]()

// Use explicitly if needed
var CachedData = h.Cached(
    5*time.Minute,
    getData,
    h.WithStore(ttlCache),
)
`

const LRUStoreExample = `
// Create an LRU cache with max 1000 items
lruCache := cache.NewLRUStore[int, UserProfile](1000)

// Use with per-key caching
var CachedUserProfile = h.CachedPerKeyT(
    30*time.Minute,
    func(userID int) (int, h.GetElementFunc) {
        return userID, func() *h.Element {
            return renderUserProfile(userID)
        }
    },
    h.WithStore(lruCache),
)
`

const MigrationBefore = `
// Existing code - continues to work without changes
var CachedDashboard = h.Cached(10*time.Minute, func() *h.Element {
    return renderDashboard()
})

var CachedUserData = h.CachedPerKeyT(15*time.Minute, func(userID string) (string, h.GetElementFunc) {
    return userID, func() *h.Element {
        return renderUserData(userID)
    }
})
`

const MigrationAfter = `
// Enhanced with custom cache store
memoryCache := cache.NewLRUStore[any, string](10000)

var CachedDashboard = h.Cached(10*time.Minute, func() *h.Element {
    return renderDashboard()
}, h.WithStore(memoryCache))

var CachedUserData = h.CachedPerKeyT(15*time.Minute, func(userID string) (string, h.GetElementFunc) {
    return userID, func() *h.Element {
        return renderUserData(userID)
    }
}, h.WithStore(memoryCache))
`

const DistributedCacheExample = `
// Initialize Redis client
redisClient := redis.NewClient(&redis.Options{
    Addr:     "redis-cluster:6379",
    Password: os.Getenv("REDIS_PASSWORD"),
    DB:       0,
})

// Create distributed cache
distributedCache := NewRedisStore[string, string](
    redisClient,
    "webapp:cache",
    30*time.Minute,
)

// Set as global default
h.DefaultCacheProvider = func() cache.Store[any, string] {
    return distributedCache
}
`

const MemoryBoundedExample = `
// Limit cache to 5000 items to prevent memory exhaustion
boundedCache := cache.NewLRUStore[string, string](5000)

// Use for user-generated content where keys might be unpredictable
var CachedSearchResults = h.CachedPerKeyT(
    5*time.Minute,
    func(query string) (string, h.GetElementFunc) {
        // Normalize and validate query to prevent cache poisoning
        normalized := normalizeSearchQuery(query)
        return normalized, func() *h.Element {
            return performSearch(normalized)
        }
    },
    h.WithStore(boundedCache),
)
`

const TieredCacheExample = `
type TieredCache[K comparable, V any] struct {
    l1 cache.Store[K, V] // Fast local cache
    l2 cache.Store[K, V] // Slower distributed cache
}

func NewTieredCache[K comparable, V any](local, distributed cache.Store[K, V]) *TieredCache[K, V] {
    return &TieredCache[K, V]{l1: local, l2: distributed}
}

func (t *TieredCache[K, V]) Get(key K) (V, bool) {
    // Check L1 first
    if val, ok := t.l1.Get(key); ok {
        return val, true
    }
    
    // Check L2
    if val, ok := t.l2.Get(key); ok {
        // Populate L1 for next time
        t.l1.Set(key, val)
        return val, true
    }
    
    var zero V
    return zero, false
}

func (t *TieredCache[K, V]) Set(key K, value V) {
    t.l1.Set(key, value)
    t.l2.Set(key, value)
}

func (t *TieredCache[K, V]) Delete(key K) {
    t.l1.Delete(key)
    t.l2.Delete(key)
}

func (t *TieredCache[K, V]) Close() error {
    if err := t.l1.Close(); err != nil {
        return err
    }
    return t.l2.Close()
}

// Usage
tieredCache := NewTieredCache(
    cache.NewLRUStore[string, string](1000),     // L1: 1k items in memory
    NewRedisStore[string, string](redis, "", 1*time.Hour), // L2: Redis
)
`
