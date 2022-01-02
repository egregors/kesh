package cache

import (
	"fmt"
	"sync"
)

// Cache is the common interface for all cache data structure implementations
type Cache interface {
	Get(key interface{}) (interface{}, error)
	Put(key interface{}, val interface{})
}

// LRUCache is Least Recently Used (LRU) cache implementation
// https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU)
type LRUCache struct {
	cache     map[interface{}]*dllNode
	cap, size int
	l         dll

	mu sync.Mutex
}

var _ Cache = (*LRUCache)(nil)

// NewLRUCache performs creating a new LRUCache instance and returns a ref to it
func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		cache: make(map[interface{}]*dllNode, capacity),
		cap:   capacity,
		size:  0,
		l:     newDll(),

		mu: sync.Mutex{},
	}
	return lru
}

// Get return value of cached object if it's possible
func (lru *LRUCache) Get(key interface{}) (interface{}, error) {
	if node, ok := lru.cache[key]; ok {
		lru.l.moveNodeToHead(node)
		return node.v, nil
	}
	return nil, fmt.Errorf("cache has not key %v", key)
}

// Put adds key-value entity into cache
func (lru *LRUCache) Put(key, val interface{}) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if node, ok := lru.cache[key]; ok {
		node.v = val
		lru.l.moveNodeToHead(node)
	} else {
		if lru.size+1 > lru.cap {
			n := lru.l.pop()
			delete(lru.cache, n.k)
			lru.size--
		}
		newNode := dllNode{
			k: key,
			v: val,
		}
		lru.cache[key] = &newNode
		lru.l.addNode(&newNode)
		lru.size++
	}
}
