package priorityqueue

import "github.com/utkarsh-pro/ugstl/pkg/heap"

type Max[T any] struct {
	heap *heap.Heap[Element[T]]
}

func NewMax[T any]() *Max[T] {
	h := heap.New(
		heap.MAX,
		func(v1, v2 Element[T]) bool {
			return v1.GetKey() < v2.GetKey()
		},
	)

	return &Max[T]{
		heap: h,
	}
}

func (m *Max[T]) Insert(data T, key int) {
	m.InsertElement(Element[T]{data: data, key: key})
}

func (m *Max[T]) InsertElement(el Element[T]) {
	m.heap.Push(el)
}

func (m *Max[T]) ExtractMax() (T, error) {
	var t T

	ele, err := m.heap.ExtractTop()
	if err != nil {
		return t, err
	}

	return ele.GetData(), nil
}

func (m *Max[T]) Max() (T, error) {
	var t T

	ele, err := m.heap.Top()
	if err != nil {
		return t, err
	}

	return ele.GetData(), nil
}
