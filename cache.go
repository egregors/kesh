package cache

// Cache is the common interface for all cache data structure implementations
type Cache interface {
	Get(key interface{}) (interface{}, error)
	Put(key interface{}, val interface{})
}
