package cache

import (
	"sync"
	"time"
)

// Cache is a key-value storage.
type Cache[K comparable, V any] struct {
	ttl  time.Duration
	data map[K]entryWithTimeout[V]
	mu   sync.Mutex
}

type entryWithTimeout[V any] struct {
	value   V
	expires time.Time // After that time, the value is useless
}

// New creates a usable cache
func New[K comparable, V any](ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:  ttl,
		data: make(map[K]entryWithTimeout[V]),
	}
}

// Read returns the associated value for a key,
// and returns a boolean to true if the key is present
func (c *Cache[K, V]) Read(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var zeroV V

	val, found := c.data[key]
	switch {
	case !found:
		return zeroV, false
	case val.expires.Before(time.Now()):
		delete(c.data, key)
		return zeroV, false
	default:
		return val.value, true
	}
}

func (c *Cache[K, V]) Upsert(key K, val V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = entryWithTimeout[V]{
		value:   val,
		expires: time.Now().Add(c.ttl),
	}

	// Do not return an error yet.
	// It can happen in the future.
	return nil
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
