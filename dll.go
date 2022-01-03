package kesh

import (
	"fmt"
	"strings"
)

// dllNode is double linked list node with arbitrary key \ value and refs to prev and next node
type dllNode struct {
	k, v       interface{}
	prev, next *dllNode
}

func (n *dllNode) remove() {
	if n.prev != nil && n.next != nil {
		prev, next := n.prev, n.next
		prev.next, next.prev = next, prev
	}
	// if node is invalid – do nothing
}

// dll is double linked list with refs to firsts and to last node of list
type dll struct {
	head, last *dllNode
}

func (l dll) String() string {
	list := []string{"head"}
	n := l.head.next

	for n.next != nil {
		list = append(list, fmt.Sprintf("%v:%v", n.k, n.v))
		n = n.next
	}
	list = append(list, "last")
	return strings.Join(list, " -> ")
}

func newDll() dll {
	l := dll{
		head: &dllNode{},
		last: &dllNode{},
	}
	l.head.next = l.last
	l.last.prev = l.head

	return l
}

func (l *dll) addNode(n *dllNode) {
	n.prev = l.head
	n.next = l.head.next
	l.head.next.prev = n
	l.head.next = n
}

func (l *dll) removeNode(n *dllNode) {
	n.remove()
}

func (l *dll) moveNodeToHead(n *dllNode) {
	l.removeNode(n)
	l.addNode(n)
}

func (l *dll) pop() *dllNode {
	toPop := l.last.prev

	if toPop != l.head {
		l.removeNode(toPop)
		return toPop
	}
	return nil
}
