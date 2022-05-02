package singly

type Node[T any] struct {
	value T
	next  *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{value: val}
}

func (n *Node[T]) GetValue() T {
	return n.value
}

func (n *Node[T]) SetValue(val T) {
	n.value = val
}

func (n *Node[T]) GetNext() *Node[T] {
	return n.next
}

func (n *Node[T]) SetNext(next *Node[T]) {
	n.next = next
}
