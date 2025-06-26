package cache

import (
	"sync"
	"testing"
	"time"
)

func TestTTLStore_SetAndGet(t *testing.T) {
	store := NewTTLStore[string, string]()
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

func TestTTLStore_Expiration(t *testing.T) {
	store := NewTTLStore[string, string]()
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

func TestTTLStore_Delete(t *testing.T) {
	store := NewTTLStore[string, string]()
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

func TestTTLStore_UpdateExisting(t *testing.T) {
	store := NewTTLStore[string, string]()
	defer store.Close()

	// Set initial value
	store.Set("key1", "value1", 100*time.Millisecond)

	// Update with new value and longer TTL
	store.Set("key1", "value2", 1*time.Hour)

	// Verify new value
	val, found := store.Get("key1")
	if !found {
		t.Error("Expected to find key1 after update")
	}
	if val != "value2" {
		t.Errorf("Expected value2, got %s", val)
	}

	// Wait for original TTL to pass
	time.Sleep(150 * time.Millisecond)

	// Should still exist with new TTL
	val, found = store.Get("key1")
	if !found {
		t.Error("Expected key1 to still exist with new TTL")
	}
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
	val, found := intStore.Get(42)
	if !found || val != "answer" {
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

	retrievedUser, found := userStore.Get("user1")
	if !found {
		t.Error("Failed to retrieve user")
	}
	if retrievedUser.ID != 1 || retrievedUser.Name != "Alice" {
		t.Error("Retrieved user data doesn't match")
	}
}
