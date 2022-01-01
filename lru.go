package main

type DLLNode struct {
	k, v       int
	prev, next *DLLNode
}

type LRUCache struct {
	cache      map[int]*DLLNode
	cap, size  int
	head, tail *DLLNode
}

func NewLRUCache(capacity int) LRUCache {
	lru := LRUCache{
		cache: make(map[int]*DLLNode),
		cap:   capacity,
		size:  0,
		head:  &DLLNode{},
		tail:  &DLLNode{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (lru *LRUCache) addNode(n *DLLNode) {
	n.prev = lru.head
	n.next = lru.head.next
	lru.head.next.prev = n
	lru.head.next = n
}

func (lru *LRUCache) removeNode(n *DLLNode) {
	prev := n.prev
	next := n.next
	prev.next = next
	next.prev = prev
}

func (lru *LRUCache) moveToHead(n *DLLNode) {
	lru.removeNode(n)
	lru.addNode(n)
}

func (lru *LRUCache) popTail() {
	res := lru.tail.prev
	lru.removeNode(res)
	delete(lru.cache, res.k)
	lru.size--
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.moveToHead(node)
		return node.v
	}
	return -1
}

func (lru *LRUCache) Put(key, value int) {
	if node, ok := lru.cache[key]; ok {
		node.v = value
		lru.moveToHead(node)
	} else {
		newNode := DLLNode{
			k: key,
			v: value,
		}
		lru.cache[key] = &newNode
		lru.addNode(&newNode)
		lru.size++

		if lru.size > lru.cap {
			lru.popTail()
		}
	}
}
