package doubly

type Node[T any] struct {
	prev  *Node[T]
	value T
	next  *Node[T]
}

func NewDNode[T any](val T) *Node[T] {
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

func (n *Node[T]) GetPrev() *Node[T] {
	return n.prev
}

func (n *Node[T]) SetPrev(prev *Node[T]) {
	n.prev = prev
}
