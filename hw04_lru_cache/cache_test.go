package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		t.Run("by capacity", func(t *testing.T) {
			c := NewCache(3)

			require.False(t, c.Set("aaa", 100))
			require.False(t, c.Set("bbb", 200))
			require.False(t, c.Set("ccc", 300))
			require.False(t, c.Set("ddd", 400))

			_, ok := c.Get("aaa")
			require.False(t, ok)

			val, ok := c.Get("bbb")
			require.True(t, ok)
			require.Equal(t, 200, val)

			val, ok = c.Get("ccc")
			require.True(t, ok)
			require.Equal(t, 300, val)

			val, ok = c.Get("ddd")
			require.True(t, ok)
			require.Equal(t, 400, val)
		})

		t.Run("by lru", func(t *testing.T) {
			c := NewCache(3)

			require.False(t, c.Set("aaa", 100))
			require.False(t, c.Set("bbb", 200))
			require.False(t, c.Set("ccc", 300))

			val, ok := c.Get("aaa")
			require.True(t, ok)
			require.Equal(t, 100, val)

			require.True(t, c.Set("bbb", 250))

			require.False(t, c.Set("ddd", 400))

			_, ok = c.Get("ccc")
			require.False(t, ok)

			val, ok = c.Get("aaa")
			require.True(t, ok)
			require.Equal(t, 100, val)

			val, ok = c.Get("bbb")
			require.True(t, ok)
			require.Equal(t, 250, val)

			val, ok = c.Get("ddd")
			require.True(t, ok)
			require.Equal(t, 400, val)
		})
	})

	t.Run("clear", func(t *testing.T) {
		c := NewCache(3)

		require.False(t, c.Set("aaa", 100))
		require.False(t, c.Set("bbb", 200))

		c.Clear()

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)

		require.False(t, c.Set("ccc", 300))
		val, ok := c.Get("ccc")
		require.True(t, ok)
		require.Equal(t, 300, val)
	})
}

const cacheIsThreadSafe = true

func TestCacheMultithreading(t *testing.T) {
	if !cacheIsThreadSafe {
		t.Skip("cache is not thread-safe yet")
	}

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
