package h

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/maddalax/htmgo/framework/h/cache"
)

func TestCached_WithDefaultStore(t *testing.T) {
	callCount := 0

	// Create a cached component
	CachedDiv := Cached(1*time.Hour, func() *Element {
		callCount++
		return Div(Text(fmt.Sprintf("Rendered %d times", callCount)))
	})

	// First render
	html1 := Render(CachedDiv())
	if callCount != 1 {
		t.Errorf("Expected 1 render, got %d", callCount)
	}

	// Second render should use cache
	html2 := Render(CachedDiv())
	if callCount != 1 {
		t.Errorf("Expected still 1 render (cached), got %d", callCount)
	}

	if html1 != html2 {
		t.Error("Expected same HTML from cache")
	}
}

func TestCached_WithCustomStore(t *testing.T) {
	// Use LRU store with small capacity
	lruStore := cache.NewLRUStore[any, string](10)
	defer lruStore.Close()

	callCount := 0

	// Create cached component with custom store
	CachedDiv := Cached(1*time.Hour, func() *Element {
		callCount++
		return Div(Text(fmt.Sprintf("Rendered %d times", callCount)))
	}, WithCacheStore(lruStore))

	// First render
	html1 := Render(CachedDiv())
	if callCount != 1 {
		t.Errorf("Expected 1 render, got %d", callCount)
	}

	// Second render should use cache
	html2 := Render(CachedDiv())
	if callCount != 1 {
		t.Errorf("Expected still 1 render (cached), got %d", callCount)
	}

	if html1 != html2 {
		t.Error("Expected same HTML from cache")
	}
}

func TestCachedPerKey_WithDefaultStore(t *testing.T) {
	renderCounts := make(map[int]int)

	// Create per-key cached component
	UserProfile := CachedPerKeyT(1*time.Hour, func(userID int) (int, GetElementFunc) {
		return userID, func() *Element {
			renderCounts[userID]++
			return Div(Text(fmt.Sprintf("User %d (rendered %d times)", userID, renderCounts[userID])))
		}
	})

	// Render for different users
	html1_user1 := Render(UserProfile(1))
	html1_user2 := Render(UserProfile(2))

	if renderCounts[1] != 1 || renderCounts[2] != 1 {
		t.Error("Expected each user to be rendered once")
	}

	// Render again - should use cache
	html2_user1 := Render(UserProfile(1))
	html2_user2 := Render(UserProfile(2))

	if renderCounts[1] != 1 || renderCounts[2] != 1 {
		t.Error("Expected renders to be cached")
	}

	if html1_user1 != html2_user1 || html1_user2 != html2_user2 {
		t.Error("Expected same HTML from cache")
	}

	// Different users should have different content
	if html1_user1 == html1_user2 {
		t.Error("Expected different content for different users")
	}
}

func TestCachedPerKey_WithLRUStore(t *testing.T) {
	// Small LRU cache that can only hold 2 items
	lruStore := cache.NewLRUStore[any, string](2)
	defer lruStore.Close()

	renderCounts := make(map[int]int)

	// Create per-key cached component with LRU store
	UserProfile := CachedPerKeyT(1*time.Hour, func(userID int) (int, GetElementFunc) {
		return userID, func() *Element {
			renderCounts[userID]++
			return Div(Text(fmt.Sprintf("User %d", userID)))
		}
	}, WithCacheStore(lruStore))

	// Render 2 users - fill cache to capacity
	Render(UserProfile(1))
	Render(UserProfile(2))

	if renderCounts[1] != 1 || renderCounts[2] != 1 {
		t.Error("Expected each user to be rendered once")
	}

	// Render user 3 - should evict user 1 (least recently used)
	Render(UserProfile(3))

	if renderCounts[3] != 1 {
		t.Error("Expected user 3 to be rendered once")
	}

	// Render user 1 again - should re-render (was evicted)
	Render(UserProfile(1))

	if renderCounts[1] != 2 {
		t.Errorf("Expected user 1 to be re-rendered after eviction, got %d renders", renderCounts[1])
	}

	// Render user 2 again - should re-render (was evicted when user 1 was added back)
	Render(UserProfile(2))

	if renderCounts[2] != 2 {
		t.Errorf("Expected user 2 to be re-rendered after eviction, got %d renders", renderCounts[2])
	}

	// At this point, cache contains users 1 and 2 (most recently used)

	// Render user 1 again - should be cached
	Render(UserProfile(1))

	if renderCounts[1] != 2 {
		t.Errorf("Expected user 1 to still be cached, got %d renders", renderCounts[1])
	}
}

func TestCachedT_WithDefaultStore(t *testing.T) {
	type Product struct {
		ID    int
		Name  string
		Price float64
	}

	renderCount := 0

	// Create cached component that takes typed data
	ProductCard := CachedT(1*time.Hour, func(p Product) *Element {
		renderCount++
		return Div(
			H3(Text(p.Name)),
			P(Text(fmt.Sprintf("$%.2f", p.Price))),
		)
	})

	product := Product{ID: 1, Name: "Widget", Price: 9.99}

	// First render
	html1 := Render(ProductCard(product))
	if renderCount != 1 {
		t.Errorf("Expected 1 render, got %d", renderCount)
	}

	// Second render should use cache
	html2 := Render(ProductCard(product))
	if renderCount != 1 {
		t.Errorf("Expected still 1 render (cached), got %d", renderCount)
	}

	if html1 != html2 {
		t.Error("Expected same HTML from cache")
	}
}

func TestCachedPerKeyT_WithCustomStore(t *testing.T) {
	type Article struct {
		ID      int
		Title   string
		Content string
	}

	ttlStore := cache.NewTTLStore[any, string]()
	defer ttlStore.Close()

	renderCounts := make(map[int]int)

	// Create per-key cached component with custom store
	ArticleView := CachedPerKeyT(1*time.Hour, func(a Article) (int, GetElementFunc) {
		return a.ID, func() *Element {
			renderCounts[a.ID]++
			return Div(
				H1(Text(a.Title)),
				P(Text(a.Content)),
			)
		}
	}, WithCacheStore(ttlStore))

	article1 := Article{ID: 1, Title: "First", Content: "Content 1"}
	article2 := Article{ID: 2, Title: "Second", Content: "Content 2"}

	// Render articles
	Render(ArticleView(article1))
	Render(ArticleView(article2))

	if renderCounts[1] != 1 || renderCounts[2] != 1 {
		t.Error("Expected each article to be rendered once")
	}

	// Render again - should use cache
	Render(ArticleView(article1))
	Render(ArticleView(article2))

	if renderCounts[1] != 1 || renderCounts[2] != 1 {
		t.Error("Expected renders to be cached")
	}
}

func TestDefaultCacheProvider_Override(t *testing.T) {
	// Save original provider
	originalProvider := DefaultCacheProvider
	defer func() {
		DefaultCacheProvider = originalProvider
	}()

	// Track which cache is used
	customCacheUsed := false

	// Override default provider
	DefaultCacheProvider = func() cache.Store[any, string] {
		customCacheUsed = true
		return cache.NewLRUStore[any, string](100)
	}

	// Create cached component without specifying store
	CachedDiv := Cached(1*time.Hour, func() *Element {
		return Div(Text("Content"))
	})

	// Render to trigger cache creation
	Render(CachedDiv())

	if !customCacheUsed {
		t.Error("Expected custom default cache provider to be used")
	}
}

func TestCachedPerKey_ConcurrentAccess(t *testing.T) {
	lruStore := cache.NewLRUStore[any, string](1000)
	defer lruStore.Close()

	UserProfile := CachedPerKeyT(1*time.Hour, func(userID int) (int, GetElementFunc) {
		return userID, func() *Element {
			// Simulate some work
			time.Sleep(10 * time.Millisecond)
			return Div(Text(fmt.Sprintf("User %d", userID)))
		}
	}, WithCacheStore(lruStore))

	const numGoroutines = 50
	const numUsers = 20

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Many goroutines accessing overlapping user IDs
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < numUsers; j++ {
				userID := j % 10 // Reuse user IDs to test cache hits
				html := Render(UserProfile(userID))

				expectedContent := fmt.Sprintf("User %d", userID)
				if !contains(html, expectedContent) {
					t.Errorf("Goroutine %d: Expected content for user %d", id, userID)
				}
			}
		}(i)
	}

	wg.Wait()
}

func TestCachedT2_MultipleParameters(t *testing.T) {
	renderCount := 0

	// Component that takes two parameters
	CombinedView := CachedT2(1*time.Hour, func(title string, count int) *Element {
		renderCount++
		return Div(
			H2(Text(title)),
			P(Text(fmt.Sprintf("Count: %d", count))),
		)
	})

	// First render
	html1 := Render(CombinedView("Test", 42))
	if renderCount != 1 {
		t.Errorf("Expected 1 render, got %d", renderCount)
	}

	// Second render with same params should use cache
	html2 := Render(CombinedView("Test", 42))
	if renderCount != 1 {
		t.Errorf("Expected still 1 render (cached), got %d", renderCount)
	}

	if html1 != html2 {
		t.Error("Expected same HTML from cache")
	}
}

func TestCachedPerKeyT3_ComplexKey(t *testing.T) {
	type CompositeKey struct {
		UserID    int
		ProductID int
		Timestamp int64
	}

	renderCount := 0

	// Component with composite key
	UserProductView := CachedPerKeyT3(1*time.Hour,
		func(userID int, productID int, timestamp int64) (CompositeKey, GetElementFunc) {
			key := CompositeKey{UserID: userID, ProductID: productID, Timestamp: timestamp}
			return key, func() *Element {
				renderCount++
				return Div(Text(fmt.Sprintf("User %d viewed product %d at %d", userID, productID, timestamp)))
			}
		},
	)

	// Render with specific combination
	ts := time.Now().Unix()
	html1 := Render(UserProductView(1, 100, ts))

	if renderCount != 1 {
		t.Errorf("Expected 1 render, got %d", renderCount)
	}

	// Same combination should use cache
	html2 := Render(UserProductView(1, 100, ts))

	if renderCount != 1 {
		t.Errorf("Expected still 1 render (cached), got %d", renderCount)
	}

	if html1 != html2 {
		t.Error("Expected same HTML from cache")
	}

	// Different combination should render again
	Render(UserProductView(1, 101, ts))

	if renderCount != 2 {
		t.Errorf("Expected 2 renders for different key, got %d", renderCount)
	}
}

func TestCached_Expiration(t *testing.T) {
	callCount := 0

	// Create cached component with short TTL
	CachedDiv := Cached(100*time.Millisecond, func() *Element {
		callCount++
		return Div(Text(fmt.Sprintf("Render %d", callCount)))
	})

	// First render
	Render(CachedDiv())
	if callCount != 1 {
		t.Errorf("Expected 1 render, got %d", callCount)
	}

	// Immediate second render should use cache
	Render(CachedDiv())
	if callCount != 1 {
		t.Errorf("Expected still 1 render (cached), got %d", callCount)
	}

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)

	// Should render again after expiration
	Render(CachedDiv())
	if callCount != 2 {
		t.Errorf("Expected 2 renders after expiration, got %d", callCount)
	}
}

func TestCachedNode_ClearCache(t *testing.T) {
	lruStore := cache.NewLRUStore[any, string](10)
	defer lruStore.Close()

	callCount := 0

	CachedDiv := Cached(1*time.Hour, func() *Element {
		callCount++
		return Div(Text("Content"))
	}, WithCacheStore(lruStore))

	// Render and cache
	element := CachedDiv()
	Render(element)

	if callCount != 1 {
		t.Errorf("Expected 1 render, got %d", callCount)
	}

	// Clear cache
	node := element.meta.(*CachedNode)
	node.ClearCache()

	// Should render again after cache clear
	Render(element)

	if callCount != 2 {
		t.Errorf("Expected 2 renders after cache clear, got %d", callCount)
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[0:len(substr)] == substr ||
		len(s) > len(substr) && contains(s[1:], substr)
}
