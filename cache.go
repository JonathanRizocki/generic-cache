package main

// Cache is a key-value storage.
type Cache[K comparable, V any] struct {
	data map[K]V
}

// New creates a usable cache
func New[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{
		data: make(map[K]V),
	}
}

// Read returns the associated value for a key,
// and returns a boolean to true if the key is present
func (c *Cache[K, V]) Read(key K) (V, bool) {
	val, found := c.data[key]
	return val, found
}

func (c *Cache[K, V]) Upsert(key K, val V) error {
	c.data[key] = val

	// Do not return an error yet.
	// It can happen in the future.
	return nil
}

func (c *Cache[K, V]) Delete(key K) {
	delete(c.data, key)
}
