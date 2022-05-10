package priorityqueue

type Element[T any] struct {
	data T
	key  int
}

func (e *Element[T]) SetKey(k int) {
	e.key = k
}

func (e *Element[T]) GetKey() int {
	return e.key
}

func (e *Element[T]) GetData() T {
	return e.data
}
