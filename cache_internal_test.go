package cache_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	cache "learngo-pockets/genericcache"
)

func TestCache(t *testing.T) {
	c := cache.New[int, string](5, time.Millisecond*1000)

	c.Upsert(5, "fünf")

	v, found := c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "fünf", v)

	c.Upsert(5, "pum")

	v, found = c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "pum", v)

	c.Upsert(3, "drei")

	v, found = c.Read(3)
	assert.True(t, found)
	assert.Equal(t, "drei", v)

	v, found = c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "pum", v)

	c.Delete(5)

	v, found = c.Read(5)
	assert.False(t, found)
	assert.Equal(t, "", v)

	v, found = c.Read(3)
	assert.True(t, found)
	assert.Equal(t, "drei", v)
}

func TestCache_Parallel_goroutines(t *testing.T) {
	c := cache.New[int, string](5, time.Millisecond*1000)

	const parallelTasks = 10
	wg := sync.WaitGroup{}
	wg.Add(parallelTasks)

	for i := 0; i < parallelTasks; i++ {
		go func(j int) {
			defer wg.Done()
			c.Upsert(4, fmt.Sprint(j))
		}(i)
	}

	wg.Wait()
}

func TestCache_Parallel(t *testing.T) {
	c := cache.New[int, string](5, time.Millisecond*1000)

	t.Run("write six", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "six")
	})

	t.Run("write kuus", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "kuus")
	})
}
