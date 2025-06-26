package cache

import (
	"sync"
	"testing"
	"time"
)

func TestLRUStore_SetAndGet(t *testing.T) {
	store := NewLRUStore[string, string](10)
	defer store.Close()

	// Test basic set and get
	store.Set("key1", "value1", 1*time.Hour)

	val, found := store.Get("key1")
	if !found {
		t.Error("Expected to find key1")
	}
	if val != "value1" {
		t.Errorf("Expected value1, got %s", val)
	}

	// Test getting non-existent key
	val, found = store.Get("nonexistent")
	if found {
		t.Error("Expected not to find nonexistent key")
	}
	if val != "" {
		t.Errorf("Expected empty string for non-existent key, got %s", val)
	}
}

func TestLRUStore_SizeLimit(t *testing.T) {
	// Create store with capacity of 3
	store := NewLRUStore[int, string](3)
	defer store.Close()

	// Add 3 items
	store.Set(1, "one", 1*time.Hour)
	store.Set(2, "two", 1*time.Hour)
	store.Set(3, "three", 1*time.Hour)

	// Verify all exist
	for i := 1; i <= 3; i++ {
		val, found := store.Get(i)
		if !found {
			t.Errorf("Expected to find key %d", i)
		}
		if val != []string{"one", "two", "three"}[i-1] {
			t.Errorf("Unexpected value for key %d: %s", i, val)
		}
	}

	// Add fourth item, should evict least recently used (key 1)
	store.Set(4, "four", 1*time.Hour)

	// Key 1 should be evicted
	_, found := store.Get(1)
	if found {
		t.Error("Expected key 1 to be evicted")
	}

	// Keys 2, 3, 4 should still exist
	for i := 2; i <= 4; i++ {
		_, found := store.Get(i)
		if !found {
			t.Errorf("Expected to find key %d", i)
		}
	}
}

func TestLRUStore_LRUBehavior(t *testing.T) {
	store := NewLRUStore[string, string](3)
	defer store.Close()

	// Add items in order
	store.Set("a", "A", 1*time.Hour)
	store.Set("b", "B", 1*time.Hour)
	store.Set("c", "C", 1*time.Hour)

	// Access "a" to make it recently used
	store.Get("a")

	// Add "d", should evict "b" (least recently used)
	store.Set("d", "D", 1*time.Hour)

	// Check what's in cache
	_, foundA := store.Get("a")
	_, foundB := store.Get("b")
	_, foundC := store.Get("c")
	_, foundD := store.Get("d")

	if !foundA {
		t.Error("Expected 'a' to still be in cache (was accessed)")
	}
	if foundB {
		t.Error("Expected 'b' to be evicted (least recently used)")
	}
	if !foundC {
		t.Error("Expected 'c' to still be in cache")
	}
	if !foundD {
		t.Error("Expected 'd' to be in cache (just added)")
	}
}

func TestLRUStore_UpdateMovesToFront(t *testing.T) {
	store := NewLRUStore[string, string](3)
	defer store.Close()

	// Fill cache
	store.Set("a", "A", 1*time.Hour)
	store.Set("b", "B", 1*time.Hour)
	store.Set("c", "C", 1*time.Hour)

	// Update "a" with new value - should move to front
	store.Set("a", "A_updated", 1*time.Hour)

	// Add new item - should evict "b" not "a"
	store.Set("d", "D", 1*time.Hour)

	val, found := store.Get("a")
	if !found {
		t.Error("Expected 'a' to still be in cache after update")
	}
	if val != "A_updated" {
		t.Errorf("Expected updated value, got %s", val)
	}

	_, found = store.Get("b")
	if found {
		t.Error("Expected 'b' to be evicted")
	}
}

func TestLRUStore_Expiration(t *testing.T) {
	store := NewLRUStore[string, string](10)
	defer store.Close()

	// Set with short TTL
	store.Set("shortlived", "value", 100*time.Millisecond)

	// Should exist immediately
	val, found := store.Get("shortlived")
	if !found {
		t.Error("Expected to find shortlived key immediately after setting")
	}
	if val != "value" {
		t.Errorf("Expected value, got %s", val)
	}

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)

	// Should be expired now
	val, found = store.Get("shortlived")
	if found {
		t.Error("Expected key to be expired")
	}
	if val != "" {
		t.Errorf("Expected empty string for expired key, got %s", val)
	}
}

func TestLRUStore_Delete(t *testing.T) {
	store := NewLRUStore[string, string](10)
	defer store.Close()

	store.Set("key1", "value1", 1*time.Hour)

	// Verify it exists
	_, found := store.Get("key1")
	if !found {
		t.Error("Expected to find key1 before deletion")
	}

	// Delete it
	store.Delete("key1")

	// Verify it's gone
	_, found = store.Get("key1")
	if found {
		t.Error("Expected key1 to be deleted")
	}

	// Delete non-existent key should not panic
	store.Delete("nonexistent")
}

func TestLRUStore_Purge(t *testing.T) {
	store := NewLRUStore[string, string](10)
	defer store.Close()

	// Add multiple items
	store.Set("key1", "value1", 1*time.Hour)
	store.Set("key2", "value2", 1*time.Hour)
	store.Set("key3", "value3", 1*time.Hour)

	// Verify they exist
	for i := 1; i <= 3; i++ {
		key := "key" + string(rune('0'+i))
		_, found := store.Get(key)
		if !found {
			t.Errorf("Expected to find %s before purge", key)
		}
	}

	// Purge all
	store.Purge()

	// Verify all are gone
	for i := 1; i <= 3; i++ {
		key := "key" + string(rune('0'+i))
		_, found := store.Get(key)
		if found {
			t.Errorf("Expected %s to be purged", key)
		}
	}
}

func TestLRUStore_ConcurrentAccess(t *testing.T) {
	// Need capacity for all unique keys: 100 goroutines * 100 operations = 10,000
	store := NewLRUStore[int, int](10000)
	defer store.Close()

	const numGoroutines = 100
	const numOperations = 100

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
				val, found := store.Get(key)
				if !found {
					t.Errorf("Goroutine %d: Expected to find key %d", id, key)
				}
				if val != key*2 {
					t.Errorf("Goroutine %d: Expected value %d, got %d", id, key*2, val)
				}
			}
		}(i)
	}

	wg.Wait()
}

func TestLRUStore_ExpiredEntriesCleanup(t *testing.T) {
	store := NewLRUStore[string, string](100)
	defer store.Close()

	// Add many short-lived entries
	for i := 0; i < 50; i++ {
		key := "key" + string(rune('0'+i))
		store.Set(key, "value", 100*time.Millisecond)
	}

	// Add some long-lived entries
	for i := 50; i < 60; i++ {
		key := "key" + string(rune('0'+i))
		store.Set(key, "value", 1*time.Hour)
	}

	// Wait for short-lived entries to expire and cleanup to run
	time.Sleep(1200 * time.Millisecond)

	// Check that expired entries are gone
	for i := 0; i < 50; i++ {
		key := "key" + string(rune('0'+i))
		_, found := store.Get(key)
		if found {
			t.Errorf("Expected expired key %s to be cleaned up", key)
		}
	}

	// Long-lived entries should still exist
	for i := 50; i < 60; i++ {
		key := "key" + string(rune('0'+i))
		_, found := store.Get(key)
		if !found {
			t.Errorf("Expected long-lived key %s to still exist", key)
		}
	}
}

func TestLRUStore_InvalidSize(t *testing.T) {
	// Test that creating store with invalid size panics
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for zero size")
		}
	}()

	NewLRUStore[string, string](0)
}

func TestLRUStore_Close(t *testing.T) {
	store := NewLRUStore[string, string](10)

	// Close should not panic
	store.Close()

	// Multiple closes should not panic
	store.Close()
	store.Close()
}

func TestLRUStore_ComplexEvictionScenario(t *testing.T) {
	store := NewLRUStore[string, string](4)
	defer store.Close()

	// Fill cache
	store.Set("a", "A", 1*time.Hour)
	store.Set("b", "B", 1*time.Hour)
	store.Set("c", "C", 1*time.Hour)
	store.Set("d", "D", 1*time.Hour)

	// Access in specific order to control LRU order
	store.Get("b") // b is most recently used
	store.Get("d") // d is second most recently used
	store.Get("a") // a is third most recently used
	// c is least recently used

	// Add two new items
	store.Set("e", "E", 1*time.Hour) // Should evict c
	store.Set("f", "F", 1*time.Hour) // Should evict the next LRU

	// Check final state
	expected := map[string]bool{
		"a": true,  // Most recently used before additions
		"b": false, // Should be evicted as second LRU
		"c": false, // First to be evicted
		"d": true,  // Second most recently used
		"e": true,  // Just added
		"f": true,  // Just added
	}

	for key, shouldExist := range expected {
		_, found := store.Get(key)
		if found != shouldExist {
			t.Errorf("Key %s: expected existence=%v, got=%v", key, shouldExist, found)
		}
	}
}
