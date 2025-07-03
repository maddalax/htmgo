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
				htmgo supports pluggable cache stores, allowing you to use any caching backend or implement custom caching strategies.
				This feature enables better control over memory usage, distributed caching support, and protection against memory exhaustion attacks.
			`),

			SubTitle("The Cache Store Interface"),
			Text(`
				All cache stores implement the following interface:
			`),
			ui.GoCodeSnippet(CacheStoreInterface),
			Text(`
				The interface is generic, supporting any comparable key type and any value type.
			`),
			Text(`
				<b>Important:</b> The <code>GetOrCompute</code> method provides <b>atomic guarantees</b>.
				When multiple goroutines request the same key simultaneously, only one will execute the compute function,
				preventing duplicate expensive operations like database queries or complex computations.
			`),

			SubTitle("Technical: The Race Condition Fix"),
			Text(`
				The previous implementation had a time-of-check to time-of-use (TOCTOU) race condition:
			`),
			Text(`
				With GetOrCompute, the entire check-compute-store operation happens atomically while holding
				the lock, eliminating the race window completely.
			`),
			Text(`
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
				The default behavior remains exactly the same, with improved concurrency guarantees.
				The framework uses the atomic GetOrCompute method internally, preventing race conditions
				that could cause duplicate renders.
			`),
			Text(`
				If you want to take advantage of custom cache stores:
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
				<b>2. Thread Safety:</b> The GetOrCompute method must be thread-safe and provide atomic guarantees.
				This means when multiple goroutines call GetOrCompute with the same key simultaneously,
				only one should execute the compute function.
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
			Text(`
				<b>6. Atomic Operations:</b> Always use GetOrCompute for cache retrieval to ensure proper
				concurrency handling and prevent cache stampedes.
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
			Text(`
				<b>Concurrency Note:</b> The GetOrCompute method eliminates race conditions that could occur
				in the previous implementation. When multiple goroutines request the same uncached key via
				GetOrCompute method simultaneously,	only one will execute the expensive render operation,
				while others wait for the result. This prevents "cache stampedes" where many goroutines
				simultaneously compute the same expensive value.
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
    // Set adds or updates an entry in the cache with the given TTL
    Set(key K, value V, ttl time.Duration)

    // GetOrCompute atomically gets an existing value or computes and stores a new value
    // This is the primary method for cache retrieval and prevents duplicate computation
    GetOrCompute(key K, compute func() V, ttl time.Duration) V

    // Delete removes an entry from the cache
    Delete(key K)

    // Purge removes all items from the cache
    Purge()

    // Close releases any resources used by the cache
    Close()
}
`

const PerComponentExample = `
// Create a custom cache store
lruCache := cache.NewLRUStore[string, string](10000) // Max 10k items

// Use it with a cached component
var CachedUserProfile = h.CachedPerKeyT(
    15*time.Minute,
    getUserProfile,
    h.WithCacheStore(lruCache), // Pass the custom store
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

func (r *RedisStore[K, V]) Set(key K, value V, ttl time.Duration) {
    ctx := context.Background()
    redisKey := fmt.Sprintf("%s:%v", r.prefix, key)

    // Serialize value
    data, err := json.Marshal(value)
    if err != nil {
        return
    }

    // Set in Redis with TTL
    r.client.Set(ctx, redisKey, data, ttl)
}

func (r *RedisStore[K, V]) GetOrCompute(key K, compute func() V, ttl time.Duration) V {
    ctx := context.Background()
    redisKey := fmt.Sprintf("%s:%v", r.prefix, key)

    // Try to get from Redis first
    data, err := r.client.Get(ctx, redisKey).Bytes()
    if err == nil {
        // Found in cache, deserialize
        var value V
        if err := json.Unmarshal(data, &value); err == nil {
            return value
        }
    }

    // Not in cache or error, compute new value
    value := compute()

    // Serialize and store
    if data, err := json.Marshal(value); err == nil {
        r.client.Set(ctx, redisKey, data, ttl)
    }

    return value
}

func (r *RedisStore[K, V]) Purge() {
    ctx := context.Background()
    // Delete all keys with our prefix
    iter := r.client.Scan(ctx, 0, r.prefix+"*", 0).Iterator()
    for iter.Next(ctx) {
        r.client.Del(ctx, iter.Val())
    }
}

func (r *RedisStore[K, V]) Delete(key K) {
    ctx := context.Background()
    redisKey := fmt.Sprintf("%s:%v", r.prefix, key)
    r.client.Del(ctx, redisKey)
}

func (r *RedisStore[K, V]) Close() {
    r.client.Close()
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
    h.WithCacheStore(redisCache),
)
`

const TTLStoreExample = `
// Create a TTL-based cache (this is the default)
ttlCache := cache.NewTTLStore[string, string]()

// Use explicitly if needed
var CachedData = h.Cached(
    5*time.Minute,
    getData,
    h.WithCacheStore(ttlCache),
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
    h.WithCacheStore(lruCache),
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
}, h.WithCacheStore(memoryCache))

var CachedUserData = h.CachedPerKeyT(15*time.Minute, func(userID string) (string, h.GetElementFunc) {
    return userID, func() *h.Element {
        return renderUserData(userID)
    }
}, h.WithCacheStore(memoryCache))
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
    h.WithCacheStore(boundedCache),
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
