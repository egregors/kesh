package kesh

import (
	"fmt"
	"sync"
)

// LRUCache is Least Recently Used (LRU) cache implementation
// https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU)
type LRUCache[K comparable, V any] struct {
	cache     map[any]*dllNode[K, V]
	cap, size int
	l         dll[K, V]

	mu sync.Mutex
}

// fixme: how should i do it now?
// var _ Cache = (*LRUCache)(nil)

// NewLRUCache performs creating a new LRUCache instance and returns a ref to it
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	// fixme: what if capacity < 0
	lru := &LRUCache[K, V]{
		cache: make(map[any]*dllNode[K, V], capacity),
		cap:   capacity,
		size:  0,
		l:     newDll[K, V](),

		mu: sync.Mutex{},
	}
	return lru
}

// Get returns value of cached object if exist either its default value
func (lru *LRUCache[K, V]) Get(key K) (val V, err error) {
	if node, ok := lru.cache[key]; ok {
		lru.l.moveNodeToHead(node)
		return node.v, nil
	}
	return val, fmt.Errorf("cache has not key %v", key)
}

// Put adds key-value entity into cache
func (lru *LRUCache[K, V]) Put(key K, val V) {
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
		newNode := dllNode[K, V]{
			k: key,
			v: val,
		}
		lru.cache[key] = &newNode
		lru.l.addNode(&newNode)
		lru.size++
	}
}
