package kesh

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLRUCache(t *testing.T) {
	c := NewLRUCache[int, int](3)
	assert.Equal(t, c.cap, 3)
}

func TestCommonLRUCache(t *testing.T) {
	lru := NewLRUCache[int, int](2)

	lru.Put(1, 1)
	lru.Put(2, 2)

	r1, err := lru.Get(1)
	require.NoError(t, err)
	assert.Equal(t, 1, r1)

	lru.Put(3, 3)
	_, err = lru.Get(2)
	require.Errorf(t, err, "cache has not key 2")

	lru.Put(4, 4)

	_, err = lru.Get(1)
	require.Errorf(t, err, "cache has not key 1")
	r3, err := lru.Get(3)
	require.NoError(t, err)
	assert.Equal(t, 3, r3)
	r4, err := lru.Get(4)
	require.NoError(t, err)
	assert.Equal(t, 4, r4)

	lru.Put(4, 42)
	r4, err = lru.Get(4)
	require.NoError(t, err)
	assert.Equal(t, 42, r4)
}

func BenchmarkLRUCache_Get_ifNotExist(b *testing.B) {
	c := NewLRUCache[int, int](1000)
	for n := 0; n < b.N; n++ {
		_, _ = c.Get(42)
	}
}

func BenchmarkLRUCache_Get_ifExist(b *testing.B) {
	c := NewLRUCache[int, string](1)
	c.Put(42, "the answer")
	for n := 0; n < b.N; n++ {
		_, _ = c.Get(42)
	}
}

func BenchmarkLRUCache_Put_inCapacity(b *testing.B) {
	c := NewLRUCache[int, string](math.MaxInt)
	for n := 0; n < b.N; n++ {
		c.Put(42, "the answer")
	}
}

func BenchmarkLRUCache_Put_outOfCapacity(b *testing.B) {
	c := NewLRUCache[int, string](1)
	for n := 0; n < b.N; n++ {
		c.Put(n, "the answer")
	}
}
