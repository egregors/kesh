package kesh

import (
	"fmt"
	"strings"
)

// dllNode is double linked list node with arbitrary key \ value and refs to prev and next node
type dllNode[K comparable, V any] struct {
	k          K
	v          V
	prev, next *dllNode[K, V]
}

func (n *dllNode[K, V]) remove() {
	if n.prev != nil && n.next != nil {
		prev, next := n.prev, n.next
		prev.next, next.prev = next, prev
	}
	// if node is invalid – do nothing
}

// dll is double linked list with refs to firsts and to last node of list
type dll[K comparable, V any] struct {
	head, last *dllNode[K, V]
}

func (l dll[K, V]) String() string {
	list := []string{"head"}
	n := l.head.next

	for n.next != nil {
		list = append(list, fmt.Sprintf("%v:%v", n.k, n.v))
		n = n.next
	}
	list = append(list, "last")
	return strings.Join(list, " -> ")
}

func newDll[K comparable, V any]() dll[K, V] {
	l := dll[K, V]{
		head: &dllNode[K, V]{},
		last: &dllNode[K, V]{},
	}
	l.head.next = l.last
	l.last.prev = l.head

	return l
}

func (l *dll[K, V]) addNode(n *dllNode[K, V]) {
	n.prev = l.head
	n.next = l.head.next
	l.head.next.prev = n
	l.head.next = n
}

func (l *dll[K, V]) removeNode(n *dllNode[K, V]) {
	n.remove()
}

func (l *dll[K, V]) moveNodeToHead(n *dllNode[K, V]) {
	l.removeNode(n)
	l.addNode(n)
}

func (l *dll[K, V]) pop() *dllNode[K, V] {
	toPop := l.last.prev
	if toPop != l.head {
		l.removeNode(toPop)
		return toPop
	}
	return nil
}
