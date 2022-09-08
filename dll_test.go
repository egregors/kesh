package kesh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dll_pop(t *testing.T) {
	l := newDll[int, string]()

	// pop empty list do nothing
	assert.Nil(t, l.pop())

	// regular pop
	l.addNode(&dllNode[int, string]{
		k: 42,
		v: "the answer",
	})
	n := l.pop()
	assert.Equal(t, 42, n.k)
	assert.Equal(t, "the answer", n.v)
}

func Test_dll_moveNodeToHead(t *testing.T) {
	l := newDll[int, int]()

	n1 := &dllNode[int, int]{k: 1, v: 1}
	n2 := &dllNode[int, int]{k: 2, v: 2}
	n3 := &dllNode[int, int]{k: 3, v: 3}

	l.addNode(n1)
	l.addNode(n2)
	l.addNode(n3)

	assert.Equal(t, l.String(), "head -> 3:3 -> 2:2 -> 1:1 -> last")
	l.moveNodeToHead(n2)
	assert.Equal(t, l.String(), "head -> 2:2 -> 3:3 -> 1:1 -> last")
	l.moveNodeToHead(n3)
	assert.Equal(t, l.String(), "head -> 3:3 -> 2:2 -> 1:1 -> last")
}

func Test_dll_removeNode(t *testing.T) {
	l := newDll[int, int]()

	n1 := &dllNode[int, int]{k: 1, v: 1}
	n2 := &dllNode[int, int]{k: 2, v: 2}
	n3 := &dllNode[int, int]{k: 3, v: 3}

	l.addNode(n1)
	l.addNode(n2)
	l.addNode(n3)

	assert.Equal(t, l.String(), "head -> 3:3 -> 2:2 -> 1:1 -> last")

	l.removeNode(n2)
	assert.Equal(t, l.String(), "head -> 3:3 -> 1:1 -> last")

	// node removing is idempotent
	l.removeNode(n2)
	assert.Equal(t, l.String(), "head -> 3:3 -> 1:1 -> last")

	l.removeNode(n1)
	assert.Equal(t, l.String(), "head -> 3:3 -> last")

	l.removeNode(n3)
	assert.Equal(t, l.String(), "head -> last")

	// removing of invalid node do nothing
	l.removeNode(&dllNode[int, int]{})
	assert.Equal(t, l.String(), "head -> last")
}

func Test_dll_addNode(t *testing.T) {
	l := newDll[int, string]()
	assert.Equal(t, l.String(), "head -> last")

	l.addNode(&dllNode[int, string]{k: 123, v: "foobar"})
	assert.Equal(t, l.String(), "head -> 123:foobar -> last")
}
