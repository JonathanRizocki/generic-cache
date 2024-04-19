package cache_test

import (
	cache "learngo-pockets/genericcache"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache_TTL(t *testing.T) {
	t.Parallel()

	c := cache.New[string, string](5, time.Millisecond*1000)
	c.Upsert("Norwegian", "Blue")

	// Check the item is there
	got, found := c.Read("Norwegian")
	assert.True(t, found)
	assert.Equal(t, "Blue", got)

	time.Sleep(time.Millisecond * 3000)

	got, found = c.Read("Norwegian")

	assert.False(t, found)
	assert.Equal(t, "", got)
}

// TestCache_MaxSize tests the maximum capacity feature of a cache.
// It checks that update items are properly requeued as "new" items,
// and that we make room by removing the most ancient item for the new ones.
func TestCache_MaxSize(t *testing.T) {
	t.Parallel()

	// Give it a TTL long enough to survive this test
	c := cache.New[int, int](3, time.Minute)

	c.Upsert(1, 1)
	c.Upsert(2, 2)
	c.Upsert(3, 3)

	got, found := c.Read(1)
	assert.True(t, found)
	assert.Equal(t, 1, got)

	// Update 1, which will no longer make it the oldest
	c.Upsert(1, 10)

	// Adding a fourth element will discard the oldest - 2 in this case.
	c.Upsert(4, 4)

	// Trying to retrieve an element that should've been discarded by now.
	got, found = c.Read(2)
	assert.False(t, found)
	assert.Equal(t, 0, got)
}
