package kesh

// Cache is the common interface for all cache data structure implementations
type Cache[K comparable, V any] interface {
	Get(key K) (V, error)
	Put(key K, val V)
}
