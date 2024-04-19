package cache

import (
	"slices"
	"sync"
	"time"
)

// Cache is a key-value storage.
type Cache[K comparable, V any] struct {
	maxSize           int
	ttl               time.Duration
	data              map[K]entryWithTimeout[V]
	mu                sync.Mutex
	chronologicalKeys []K
}

type entryWithTimeout[V any] struct {
	value   V
	expires time.Time // After that time, the value is useless
}

// New creates a usable cache
func New[K comparable, V any](maxSize int, ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:               ttl,
		data:              make(map[K]entryWithTimeout[V]),
		chronologicalKeys: make([]K, 0, maxSize),
		maxSize:           maxSize,
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
		c.deleteKeyValue(key)
		return zeroV, false
	default:
		return val.value, true
	}
}

func (c *Cache[K, V]) Upsert(key K, val V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, alreadyPresent := c.data[key]

	switch {
	case alreadyPresent:
		c.deleteKeyValue(key)
	case len(c.data) == c.maxSize:
		c.deleteKeyValue(c.chronologicalKeys[0])
	}

	c.addKeyValue(key, val)

	// Do not return an error yet.
	// It can happen in the future.
	return nil
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.deleteKeyValue(key)
}

func (c *Cache[K, V]) addKeyValue(key K, val V) {
	c.data[key] = entryWithTimeout[V]{
		value:   val,
		expires: time.Now().Add(c.ttl),
	}
	c.chronologicalKeys = append(c.chronologicalKeys, key)
}

func (c *Cache[K, V]) deleteKeyValue(key K) {
	c.chronologicalKeys = slices.DeleteFunc(c.chronologicalKeys, func(k K) bool { return k == key })
	delete(c.data, key)
}
