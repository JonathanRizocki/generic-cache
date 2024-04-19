package cache_test

import (
	cache "learngo-pockets/genericcache"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache_TTL(t *testing.T) {
	t.Parallel()

	c := cache.New[string, string](1 * time.Millisecond * 1000)
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
