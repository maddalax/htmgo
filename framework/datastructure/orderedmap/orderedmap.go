package orderedmap

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// Map is a generic data structure that maintains the order of keys.
type Map[K comparable, V any] struct {
	keys   []K
	values map[K]V
}

func (om *Map[K, V]) Each(cb func(key K, value V)) {
	for _, key := range om.keys {
		cb(key, om.values[key])
	}
}

// Entries returns the key-value pairs in the order they were added.
func (om *Map[K, V]) Entries() []Entry[K, V] {
	entries := make([]Entry[K, V], len(om.keys))
	for i, key := range om.keys {
		entries[i] = Entry[K, V]{
			Key:   key,
			Value: om.values[key],
		}
	}
	return entries
}

// New creates a new Map.
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		keys:   []K{},
		values: make(map[K]V),
	}
}

// Set adds or updates a key-value pair in the Map.
func (om *Map[K, V]) Set(key K, value V) {
	// Check if the key already exists
	if _, exists := om.values[key]; !exists {
		om.keys = append(om.keys, key) // Append key to the keys slice if it's a new key
	}
	om.values[key] = value
}

// Get retrieves a value by key.
func (om *Map[K, V]) Get(key K) (V, bool) {
	value, exists := om.values[key]
	return value, exists
}

// Keys returns the keys in the order they were added.
func (om *Map[K, V]) Keys() []K {
	return om.keys
}

// Values returns the values in the order of their keys.
func (om *Map[K, V]) Values() []V {
	values := make([]V, len(om.keys))
	for i, key := range om.keys {
		values[i] = om.values[key]
	}

	return values
}

// Delete removes a key-value pair from the Map.
func (om *Map[K, V]) Delete(key K) {
	if _, exists := om.values[key]; exists {
		// Remove the key from the map
		delete(om.values, key)

		// Remove the key from the keys slice
		for i, k := range om.keys {
			if k == key {
				om.keys = append(om.keys[:i], om.keys[i+1:]...)
				break
			}
		}
	}
}
