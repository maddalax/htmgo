package cache

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestTTLStore_SetAndGet(t *testing.T) {
	store := NewTTLStore[string, string]()
	defer store.Close()

	// Test basic set and get
	store.Set("key1", "value1", 1*time.Hour)

	val := store.GetOrCompute("key1", func() string {
		t.Error("Should not compute for existing key")
		return "should-not-compute"
	}, 1*time.Hour)
	if val != "value1" {
		t.Errorf("Expected value1, got %s", val)
	}

	// Test getting non-existent key
	computeCalled := false
	val = store.GetOrCompute("nonexistent", func() string {
		computeCalled = true
		return "computed-value"
	}, 1*time.Hour)
	if !computeCalled {
		t.Error("Expected compute function to be called for non-existent key")
	}
	if val != "computed-value" {
		t.Errorf("Expected computed-value for non-existent key, got %s", val)
	}
}

func TestTTLStore_Expiration(t *testing.T) {
	store := NewTTLStore[string, string]()
	defer store.Close()

	// Set with short TTL
	store.Set("shortlived", "value", 100*time.Millisecond)

	// Should exist immediately
	val := store.GetOrCompute("shortlived", func() string {
		t.Error("Should not compute for existing key")
		return "should-not-compute"
	}, 100*time.Millisecond)
	if val != "value" {
		t.Errorf("Expected value, got %s", val)
	}

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)

	// Should be expired now
	computeCalled := false
	val = store.GetOrCompute("shortlived", func() string {
		computeCalled = true
		return "recomputed-after-expiry"
	}, 100*time.Millisecond)
	if !computeCalled {
		t.Error("Expected compute function to be called for expired key")
	}
	if val != "recomputed-after-expiry" {
		t.Errorf("Expected recomputed value for expired key, got %s", val)
	}
}

func TestTTLStore_Delete(t *testing.T) {
	store := NewTTLStore[string, string]()
	defer store.Close()

	store.Set("key1", "value1", 1*time.Hour)

	// Verify it exists
	val := store.GetOrCompute("key1", func() string {
		t.Error("Should not compute for existing key")
		return "should-not-compute"
	}, 1*time.Hour)
	if val != "value1" {
		t.Errorf("Expected value1, got %s", val)
	}

	// Delete it
	store.Delete("key1")

	// Verify it's gone
	computeCalled := false
	val = store.GetOrCompute("key1", func() string {
		computeCalled = true
		return "recomputed-after-delete"
	}, 1*time.Hour)
	if !computeCalled {
		t.Error("Expected compute function to be called after deletion")
	}
	if val != "recomputed-after-delete" {
		t.Errorf("Expected recomputed value after deletion, got %s", val)
	}

	// Delete non-existent key should not panic
	store.Delete("nonexistent")
}

func TestTTLStore_Purge(t *testing.T) {
	store := NewTTLStore[string, string]()
	defer store.Close()

	// Add multiple items
	store.Set("key1", "value1", 1*time.Hour)
	store.Set("key2", "value2", 1*time.Hour)
	store.Set("key3", "value3", 1*time.Hour)

	// Verify they exist
	for i := 1; i <= 3; i++ {
		key := "key" + string(rune('0'+i))
		val := store.GetOrCompute(key, func() string {
			t.Errorf("Should not compute for existing key %s", key)
			return "should-not-compute"
		}, 1*time.Hour)
		expectedVal := "value" + string(rune('0'+i))
		if val != expectedVal {
			t.Errorf("Expected to find %s with value %s, got %s", key, expectedVal, val)
		}
	}

	// Purge all
	store.Purge()

	// Verify all are gone
	for i := 1; i <= 3; i++ {
		key := "key" + string(rune('0'+i))
		computeCalled := false
		store.GetOrCompute(key, func() string {
			computeCalled = true
			return "recomputed-after-purge"
		}, 1*time.Hour)
		if !computeCalled {
			t.Errorf("Expected %s to be purged and recomputed", key)
		}
	}
}

func TestTTLStore_ConcurrentAccess(t *testing.T) {
	store := NewTTLStore[int, int]()
	defer store.Close()

	const numGoroutines = 100
	const numOperations = 1000

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Concurrent writes and reads
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				key := (id * numOperations) + j
				store.Set(key, key*2, 1*time.Hour)

				// Immediately read it back
				val := store.GetOrCompute(key, func() int {
					t.Errorf("Goroutine %d: Should not compute for just-set key %d", id, key)
					return -1
				}, 1*time.Hour)
				if val != key*2 {
					t.Errorf("Goroutine %d: Expected value %d, got %d", id, key*2, val)
				}
			}
		}(i)
	}

	wg.Wait()
}

func TestTTLStore_UpdateExisting(t *testing.T) {
	store := NewTTLStore[string, string]()
	defer store.Close()

	// Set initial value
	store.Set("key1", "value1", 100*time.Millisecond)

	// Update with new value and longer TTL
	store.Set("key1", "value2", 1*time.Hour)

	// Verify new value
	val := store.GetOrCompute("key1", func() string {
		t.Error("Should not compute for existing key")
		return "should-not-compute"
	}, 1*time.Hour)
	if val != "value2" {
		t.Errorf("Expected value2, got %s", val)
	}

	// Wait for original TTL to pass
	time.Sleep(150 * time.Millisecond)

	// Should still exist with new TTL
	val = store.GetOrCompute("key1", func() string {
		t.Error("Should not compute for key with new TTL")
		return "should-not-compute"
	}, 1*time.Hour)
	if val != "value2" {
		t.Errorf("Expected value2, got %s", val)
	}
}

func TestTTLStore_CleanupGoroutine(t *testing.T) {
	// This test verifies that expired entries are cleaned up automatically
	store := NewTTLStore[string, string]()
	defer store.Close()

	// Add many short-lived entries
	for i := 0; i < 100; i++ {
		key := "key" + string(rune('0'+i))
		store.Set(key, "value", 100*time.Millisecond)
	}

	// Cast to access internal state for testing
	ttlStore := store.(*TTLStore[string, string])

	// Check initial count
	ttlStore.mutex.RLock()
	initialCount := len(ttlStore.cache)
	ttlStore.mutex.RUnlock()

	if initialCount != 100 {
		t.Errorf("Expected 100 entries initially, got %d", initialCount)
	}

	// Wait for expiration and cleanup cycle
	// In test mode, cleanup runs every second
	time.Sleep(1200 * time.Millisecond)

	// Check that entries were cleaned up
	ttlStore.mutex.RLock()
	finalCount := len(ttlStore.cache)
	ttlStore.mutex.RUnlock()

	if finalCount != 0 {
		t.Errorf("Expected 0 entries after cleanup, got %d", finalCount)
	}
}

func TestTTLStore_Close(t *testing.T) {
	store := NewTTLStore[string, string]()

	// Close should not panic
	store.Close()

	// Multiple closes should not panic
	store.Close()
	store.Close()
}

func TestTTLStore_DifferentTypes(t *testing.T) {
	// Test with different key and value types
	intStore := NewTTLStore[int, string]()
	defer intStore.Close()

	intStore.Set(42, "answer", 1*time.Hour)
	val := intStore.GetOrCompute(42, func() string {
		t.Error("Should not compute for existing key")
		return "should-not-compute"
	}, 1*time.Hour)
	if val != "answer" {
		t.Error("Failed with int key")
	}

	// Test with struct values
	type User struct {
		ID   int
		Name string
	}

	userStore := NewTTLStore[string, User]()
	defer userStore.Close()

	user := User{ID: 1, Name: "Alice"}
	userStore.Set("user1", user, 1*time.Hour)

	retrievedUser := userStore.GetOrCompute("user1", func() User {
		t.Error("Should not compute for existing user")
		return User{}
	}, 1*time.Hour)
	if retrievedUser.ID != 1 || retrievedUser.Name != "Alice" {
		t.Error("Retrieved user data doesn't match")
	}
}

func TestTTLStore_GetOrCompute(t *testing.T) {
	store := NewTTLStore[string, string]()
	defer store.Close()

	computeCount := 0

	// Test computing when not in cache
	result := store.GetOrCompute("key1", func() string {
		computeCount++
		return "computed-value"
	}, 1*time.Hour)

	if result != "computed-value" {
		t.Errorf("Expected computed-value, got %s", result)
	}
	if computeCount != 1 {
		t.Errorf("Expected compute to be called once, called %d times", computeCount)
	}

	// Test returning cached value
	result = store.GetOrCompute("key1", func() string {
		computeCount++
		return "should-not-compute"
	}, 1*time.Hour)

	if result != "computed-value" {
		t.Errorf("Expected cached value, got %s", result)
	}
	if computeCount != 1 {
		t.Errorf("Expected compute to not be called again, total calls: %d", computeCount)
	}
}

func TestTTLStore_GetOrCompute_Expiration(t *testing.T) {
	store := NewTTLStore[string, string]()
	defer store.Close()

	computeCount := 0

	// Set with short TTL
	result := store.GetOrCompute("shortlived", func() string {
		computeCount++
		return "value1"
	}, 100*time.Millisecond)

	if result != "value1" {
		t.Errorf("Expected value1, got %s", result)
	}
	if computeCount != 1 {
		t.Errorf("Expected 1 compute, got %d", computeCount)
	}

	// Should return cached value immediately
	result = store.GetOrCompute("shortlived", func() string {
		computeCount++
		return "value2"
	}, 100*time.Millisecond)

	if result != "value1" {
		t.Errorf("Expected cached value1, got %s", result)
	}
	if computeCount != 1 {
		t.Errorf("Expected still 1 compute, got %d", computeCount)
	}

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)

	// Should compute new value after expiration
	result = store.GetOrCompute("shortlived", func() string {
		computeCount++
		return "value2"
	}, 100*time.Millisecond)

	if result != "value2" {
		t.Errorf("Expected new value2, got %s", result)
	}
	if computeCount != 2 {
		t.Errorf("Expected 2 computes after expiration, got %d", computeCount)
	}
}

func TestTTLStore_GetOrCompute_Concurrent(t *testing.T) {
	store := NewTTLStore[string, string]()
	defer store.Close()

	var computeCount int32
	const numGoroutines = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Launch many goroutines trying to compute the same key
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			result := store.GetOrCompute("shared-key", func() string {
				// Increment atomically to count calls
				atomic.AddInt32(&computeCount, 1)
				// Simulate some work
				time.Sleep(10 * time.Millisecond)
				return "shared-value"
			}, 1*time.Hour)

			if result != "shared-value" {
				t.Errorf("Goroutine %d: Expected shared-value, got %s", id, result)
			}
		}(i)
	}

	wg.Wait()

	// Only one goroutine should have computed the value
	if computeCount != 1 {
		t.Errorf("Expected exactly 1 compute for concurrent access, got %d", computeCount)
	}
}

func TestTTLStore_GetOrCompute_MultipleKeys(t *testing.T) {
	store := NewTTLStore[int, int]()
	defer store.Close()

	computeCounts := make(map[int]int)
	var mu sync.Mutex

	// Test multiple different keys
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ { // Access each key 3 times
			result := store.GetOrCompute(i, func() int {
				mu.Lock()
				computeCounts[i]++
				mu.Unlock()
				return i * 10
			}, 1*time.Hour)

			if result != i*10 {
				t.Errorf("Expected %d, got %d", i*10, result)
			}
		}
	}

	// Each key should be computed exactly once
	for i := 0; i < 10; i++ {
		if computeCounts[i] != 1 {
			t.Errorf("Key %d: Expected 1 compute, got %d", i, computeCounts[i])
		}
	}
}
