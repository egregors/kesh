// Package kesh is collection of some cache implementation.
//
// All thus interfaces are only for documentation purpose, do not use it in consumer code.
// Create your oun interfaces.
package kesh

// Cache is the common interface for all cache data structure implementations.
type Cache[K comparable, V any] interface {
	Get(key K) (V, error)
	Put(key K, val V)
}
