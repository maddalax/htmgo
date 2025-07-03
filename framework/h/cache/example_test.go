package cache_test

import (
	"fmt"
	"sync"
	"time"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/h/cache"
)

// Example demonstrates basic caching with the default TTL store
func ExampleCached() {
	renderCount := 0

	// Create a cached component that expires after 5 minutes
	CachedHeader := h.Cached(5*time.Minute, func() *h.Element {
		renderCount++
		return h.Header(
			h.H1(h.Text("Welcome to our site")),
			h.P(h.Text(fmt.Sprintf("Rendered %d times", renderCount))),
		)
	})

	// First render - will execute the function
	html1 := h.Render(CachedHeader())
	fmt.Println("Render count:", renderCount)

	// Second render - will use cached HTML
	html2 := h.Render(CachedHeader())
	fmt.Println("Render count:", renderCount)
	fmt.Println("Same HTML:", html1 == html2)

	// Output:
	// Render count: 1
	// Render count: 1
	// Same HTML: true
}

// Example demonstrates per-key caching for user-specific content
func ExampleCachedPerKeyT() {
	type User struct {
		ID   int
		Name string
	}

	renderCounts := make(map[int]int)

	// Create a per-user cached component
	UserProfile := h.CachedPerKeyT(15*time.Minute, func(user User) (int, h.GetElementFunc) {
		// Use user ID as the cache key
		return user.ID, func() *h.Element {
			renderCounts[user.ID]++
			return h.Div(
				h.Class("user-profile"),
				h.H2(h.Text(user.Name)),
				h.P(h.Text(fmt.Sprintf("User ID: %d", user.ID))),
			)
		}
	})

	alice := User{ID: 1, Name: "Alice"}
	bob := User{ID: 2, Name: "Bob"}

	// Render Alice's profile - will execute
	h.Render(UserProfile(alice))
	fmt.Printf("Alice render count: %d\n", renderCounts[1])

	// Render Bob's profile - will execute
	h.Render(UserProfile(bob))
	fmt.Printf("Bob render count: %d\n", renderCounts[2])

	// Render Alice's profile again - will use cache
	h.Render(UserProfile(alice))
	fmt.Printf("Alice render count after cache hit: %d\n", renderCounts[1])

	// Output:
	// Alice render count: 1
	// Bob render count: 1
	// Alice render count after cache hit: 1
}

// Example demonstrates using a memory-bounded LRU cache
func ExampleWithStore_lru() {
	// Create an LRU cache that holds maximum 1000 items
	lruStore := cache.NewLRUStore[any, string](1000)
	defer lruStore.Close()

	renderCount := 0

	// Use the LRU cache for a component
	ProductCard := h.CachedPerKeyT(1*time.Hour,
		func(productID int) (int, h.GetElementFunc) {
			return productID, func() *h.Element {
				renderCount++
				// Simulate fetching product data
				return h.Div(
					h.H3(h.Text(fmt.Sprintf("Product #%d", productID))),
					h.P(h.Text("$99.99")),
				)
			}
		},
		h.WithStore(lruStore), // Use custom cache store
	)

	// Render many products
	for i := 0; i < 1500; i++ {
		h.Render(ProductCard(i))
	}

	// Due to LRU eviction, only 1000 items are cached
	// Earlier items (0-499) were evicted
	fmt.Printf("Total renders: %d\n", renderCount)
	fmt.Printf("Expected renders: %d (due to LRU eviction)\n", 1500)

	// Accessing an evicted item will cause a re-render
	h.Render(ProductCard(0))
	fmt.Printf("After accessing evicted item: %d\n", renderCount)

	// Output:
	// Total renders: 1500
	// Expected renders: 1500 (due to LRU eviction)
	// After accessing evicted item: 1501
}

// MockDistributedCache simulates a distributed cache like Redis
type MockDistributedCache struct {
	data  map[string]string
	mutex sync.RWMutex
}

// DistributedCacheAdapter makes MockDistributedCache compatible with cache.Store interface
type DistributedCacheAdapter struct {
	cache *MockDistributedCache
}

func (a *DistributedCacheAdapter) Set(key any, value string, ttl time.Duration) {
	a.cache.mutex.Lock()
	defer a.cache.mutex.Unlock()
	// In a real implementation, you'd set TTL in Redis
	keyStr := fmt.Sprintf("htmgo:%v", key)
	a.cache.data[keyStr] = value
}



func (a *DistributedCacheAdapter) Delete(key any) {
	a.cache.mutex.Lock()
	defer a.cache.mutex.Unlock()
	keyStr := fmt.Sprintf("htmgo:%v", key)
	delete(a.cache.data, keyStr)
}

func (a *DistributedCacheAdapter) Purge() {
	a.cache.mutex.Lock()
	defer a.cache.mutex.Unlock()
	a.cache.data = make(map[string]string)
}

func (a *DistributedCacheAdapter) Close() {
	// Clean up connections in real implementation
}

func (a *DistributedCacheAdapter) GetOrCompute(key any, compute func() string, ttl time.Duration) string {
	a.cache.mutex.Lock()
	defer a.cache.mutex.Unlock()
	
	keyStr := fmt.Sprintf("htmgo:%v", key)
	
	// Check if exists
	if val, ok := a.cache.data[keyStr]; ok {
		return val
	}
	
	// Compute and store
	value := compute()
	a.cache.data[keyStr] = value
	// In a real implementation, you'd also set TTL in Redis
	
	return value
}

// Example demonstrates creating a custom cache adapter
func ExampleDistributedCacheAdapter() {

	// Create the distributed cache
	distCache := &MockDistributedCache{
		data: make(map[string]string),
	}
	adapter := &DistributedCacheAdapter{cache: distCache}

	// Use it with a cached component
	SharedComponent := h.Cached(10*time.Minute, func() *h.Element {
		return h.Div(h.Text("Shared across all servers"))
	}, h.WithStore(adapter))

	html := h.Render(SharedComponent())
	fmt.Printf("Cached in distributed store: %v\n", len(distCache.data) > 0)
	fmt.Printf("HTML length: %d\n", len(html))

	// Output:
	// Cached in distributed store: true
	// HTML length: 36
}

// Example demonstrates overriding the default cache provider globally
func ExampleDefaultCacheProvider() {
	// Save the original provider to restore it later
	originalProvider := h.DefaultCacheProvider
	defer func() {
		h.DefaultCacheProvider = originalProvider
	}()

	// Override the default to use LRU for all cached components
	h.DefaultCacheProvider = func() cache.Store[any, string] {
		// All cached components will use 10,000 item LRU cache by default
		return cache.NewLRUStore[any, string](10_000)
	}

	// Now all cached components use LRU by default
	renderCount := 0
	AutoLRUComponent := h.Cached(1*time.Hour, func() *h.Element {
		renderCount++
		return h.Div(h.Text("Using LRU by default"))
	})

	h.Render(AutoLRUComponent())
	fmt.Printf("Render count: %d\n", renderCount)

	// Output:
	// Render count: 1
}

// Example demonstrates caching with complex keys
func ExampleCachedPerKeyT3() {
	type FilterOptions struct {
		Category string
		MinPrice float64
		MaxPrice float64
	}

	renderCount := 0

	// Cache filtered product lists with composite keys
	FilteredProducts := h.CachedPerKeyT3(30*time.Minute,
		func(category string, minPrice, maxPrice float64) (FilterOptions, h.GetElementFunc) {
			// Create composite key from all parameters
			key := FilterOptions{
				Category: category,
				MinPrice: minPrice,
				MaxPrice: maxPrice,
			}
			return key, func() *h.Element {
				renderCount++
				// Simulate database query with filters
				return h.Div(
					h.H3(h.Text(fmt.Sprintf("Products in %s", category))),
					h.P(h.Text(fmt.Sprintf("Price range: $%.2f - $%.2f", minPrice, maxPrice))),
					h.Ul(
						h.Li(h.Text("Product 1")),
						h.Li(h.Text("Product 2")),
						h.Li(h.Text("Product 3")),
					),
				)
			}
		},
	)

	// First query - will render
	h.Render(FilteredProducts("Electronics", 100.0, 500.0))
	fmt.Printf("Render count: %d\n", renderCount)

	// Same query - will use cache
	h.Render(FilteredProducts("Electronics", 100.0, 500.0))
	fmt.Printf("Render count after cache hit: %d\n", renderCount)

	// Different query - will render
	h.Render(FilteredProducts("Electronics", 200.0, 600.0))
	fmt.Printf("Render count after new query: %d\n", renderCount)

	// Output:
	// Render count: 1
	// Render count after cache hit: 1
	// Render count after new query: 2
}

// Example demonstrates cache expiration and refresh
func ExampleCached_expiration() {
	renderCount := 0
	now := time.Now()

	// Cache with very short TTL for demonstration
	TimeSensitive := h.Cached(100*time.Millisecond, func() *h.Element {
		renderCount++
		return h.Div(
			h.Text(fmt.Sprintf("Generated at: %s (render #%d)",
				now.Format("15:04:05"), renderCount)),
		)
	})

	// First render
	h.Render(TimeSensitive())
	fmt.Printf("Render count: %d\n", renderCount)

	// Immediate second render - uses cache
	h.Render(TimeSensitive())
	fmt.Printf("Render count (cached): %d\n", renderCount)

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)

	// Render after expiration - will re-execute
	h.Render(TimeSensitive())
	fmt.Printf("Render count (after expiration): %d\n", renderCount)

	// Output:
	// Render count: 1
	// Render count (cached): 1
	// Render count (after expiration): 2
}
