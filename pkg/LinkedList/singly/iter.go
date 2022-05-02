package singly

type Iter[T any] struct {
	list *List[T]

	curr *Node[T]
}

func NewIter[T any](list *List[T]) *Iter[T] {
	i := Iter[T]{list: list}
	i.curr = list.head

	return &i
}

func (i *Iter[T]) Next() (*Node[T], bool) {
	var val *Node[T]

	if i.curr == nil {
		return val, false
	}

	val = i.curr
	i.curr = i.curr.GetNext()

	return val, true
}
