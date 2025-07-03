package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/h/cache"
)

// This example demonstrates the atomic guarantees of GetOrCompute,
// showing how it prevents duplicate expensive computations when
// multiple goroutines request the same uncached key simultaneously.

func main() {
	fmt.Println("=== Atomic Cache Example ===")

	// Demonstrate the problem without atomic guarantees
	demonstrateProblem()

	fmt.Println("\n=== Now with GetOrCompute atomic guarantees ===")

	// Show the solution with GetOrCompute
	demonstrateSolution()

}

// demonstrateProblem shows what happens without atomic guarantees
func demonstrateProblem() {
	fmt.Println("Without atomic guarantees (simulated):")
	fmt.Println("Multiple goroutines checking cache and computing...")

	var computeCount int32
	var wg sync.WaitGroup

	// Simulate 10 goroutines trying to get the same uncached value
	for i := range 10 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Simulate checking cache (not found)
			time.Sleep(time.Millisecond) // Small delay to increase collision chance

			// All goroutines think the value is not cached
			// so they all compute it
			atomic.AddInt32(&computeCount, 1)
			fmt.Printf("Goroutine %d: Computing expensive value...\n", id)

			// Simulate expensive computation
			time.Sleep(50 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	fmt.Printf("\nResult: Computed %d times (wasteful!)\n", computeCount)
}

// demonstrateSolution shows how GetOrCompute solves the problem
func demonstrateSolution() {
	// Create a cache store
	store := cache.NewTTLStore[string, string]()
	defer store.Close()

	var computeCount int32
	var wg sync.WaitGroup

	fmt.Println("With GetOrCompute atomic guarantees:")
	fmt.Println("Multiple goroutines requesting the same key...")

	startTime := time.Now()

	// Launch 10 goroutines trying to get the same value
	for i := range 10 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// All goroutines call GetOrCompute at the same time
			result := store.GetOrCompute("expensive-key", func() string {
				// Only ONE goroutine will execute this function
				count := atomic.AddInt32(&computeCount, 1)
				fmt.Printf("Goroutine %d: Computing expensive value (computation #%d)\n", id, count)

				// Simulate expensive computation
				time.Sleep(50 * time.Millisecond)

				return fmt.Sprintf("Expensive result computed by goroutine %d", id)
			}, 1*time.Hour)

			fmt.Printf("Goroutine %d: Got result: %s\n", id, result)
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(startTime)

	fmt.Printf("\nResult: Computed only %d time (efficient!)\n", computeCount)
	fmt.Printf("Total time: %v (vs ~500ms if all computed)\n", elapsed)
}

// Example with htmgo cached components
func ExampleCachedComponent() {
	fmt.Println("\n=== Real-world htmgo Example ===")

	var renderCount int32

	// Create a cached component that simulates fetching user data
	UserProfile := h.CachedPerKeyT(5*time.Minute, func(userID int) (int, h.GetElementFunc) {
		return userID, func() *h.Element {
			count := atomic.AddInt32(&renderCount, 1)
			fmt.Printf("Fetching and rendering user %d (render #%d)\n", userID, count)

			// Simulate database query
			time.Sleep(100 * time.Millisecond)

			return h.Div(
				h.H2(h.Text(fmt.Sprintf("User Profile #%d", userID))),
				h.P(h.Text("This was expensive to compute!")),
			)
		}
	})

	// Simulate multiple concurrent requests for the same user
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func(requestID int) {
			defer wg.Done()

			// All requests are for user 123
			html := h.Render(UserProfile(123))
			fmt.Printf("Request %d: Received %d bytes of HTML\n", requestID, len(html))
		}(i)
	}

	wg.Wait()
	fmt.Printf("\nTotal renders: %d (only one, despite 5 concurrent requests!)\n", renderCount)
}

// Example showing cache stampede prevention
func ExampleCacheStampedePrevention() {
	fmt.Println("\n=== Cache Stampede Prevention ===")

	store := cache.NewLRUStore[string, string](100)
	defer store.Close()

	var dbQueries int32

	// Simulate a popular cache key expiring
	fetchPopularData := func(key string) string {
		return store.GetOrCompute(key, func() string {
			queries := atomic.AddInt32(&dbQueries, 1)
			fmt.Printf("Database query #%d for key: %s\n", queries, key)

			// Simulate slow database query
			time.Sleep(200 * time.Millisecond)

			return fmt.Sprintf("Popular data for %s", key)
		}, 100*time.Millisecond) // Short TTL to simulate expiration
	}

	// First, populate the cache
	_ = fetchPopularData("trending-posts")
	fmt.Println("Cache populated")

	// Wait for it to expire
	time.Sleep(150 * time.Millisecond)
	fmt.Println("\nCache expired, simulating traffic spike...")

	// Simulate 20 concurrent requests right after expiration
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			data := fetchPopularData("trending-posts")
			fmt.Printf("Request %d: Got data: %s\n", id, data)
		}(i)
	}

	wg.Wait()
	fmt.Printf("\nTotal database queries: %d (prevented 19 redundant queries!)\n", dbQueries)
}
