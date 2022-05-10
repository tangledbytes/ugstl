package heap

import "github.com/utkarsh-pro/ugstl/pkg/types"

type Type int

const (
	MIN Type = iota
	MAX
)

type Heap[T any] struct {
	array  Store[T]
	size   int
	compFn types.Less[T]
}

// New takes the type of the Heap (MIN/MAX) and a less function
// which must return true if v1 < v2, returns a pointer to the Heap
func New[T any](typ Type, less types.Less[T]) *Heap[T] {
	h := &Heap[T]{
		array: Store[T]{},
		size:  0,
	}

	switch typ {
	case MAX:
		h.compFn = less
	case MIN:
		h.compFn = func(v1, v2 T) bool {
			return less(v2, v1)
		}
	default:
		panic(ErrInvalidHeapConfig)
	}

	return h
}

// NewFromArray takes the array, type of the Heap (MIN/MAX) and a less function
// which must return true if v1 < v2, returns a pointer to the Heap
func NewFromArray[T any](arr []T, typ Type, less types.Less[T]) *Heap[T] {
	h := New(typ, less)
	h.array = Store[T](arr)

	h.Fix()

	return h
}

// Fix will attempt to rebuild the Heap
//
// This method must be called if manual alteration has been
// performed to the underlying array
func (h *Heap[T]) Fix() {
	h.size = len(h.array)
	for i := len(h.array)/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *Heap[T]) Set(idx int, el T) error {
	if idx < 0 || idx > h.size {
		return ErrIndexOutOfBound
	}

	orig := h.array[idx]

	if !h.compFn(orig, el) {
		return ErrInvalidElement
	}

	h.array[idx] = el

	for idx > 0 && h.compFn(h.array[parentIdx(idx)], h.array[idx]) {
		h.array[parentIdx(idx)], h.array[idx] = h.array[idx], h.array[parentIdx(idx)]
		idx = parentIdx(idx)
	}

	return nil
}

// Push adds the element to the heap
func (h *Heap[T]) Push(el T) {
	h.array = append(h.array, el)
	h.size++

	for {
		p := parentIdx(h.size - 1)

		if h.compFn(h.array[p], h.array[h.size-1]) {
			h.array[p], h.array[h.size-1] = h.array[h.size-1], h.array[p]
		} else {
			break
		}
	}
}

// Top returns the top element of the heap
func (h *Heap[T]) Top() (T, error) {
	var t T

	if h.size < 1 {
		return t, ErrHeapUnderflow
	}

	t = h.array[0]
	return t, nil
}

func (h *Heap[T]) ExtractTop() (T, error) {
	var t T

	if h.size < 1 {
		return t, ErrHeapUnderflow
	}

	h.array[0], h.array[h.size-1] = h.array[h.size-1], h.array[0]
	t = h.array[h.size-1]

	h.array = h.array[:h.size-1]
	h.size--
	h.heapify(0)

	return t, nil
}

// heapify takes the index of the element and heapifies the
// the heap on that index
func (h *Heap[T]) heapify(idx int) {
	l := leftIdx(idx)
	r := rightIdx(idx)

	target := idx

	if l < h.size && h.compFn(h.array[target], h.array[l]) {
		target = l
	}
	if r < h.size && h.compFn(h.array[target], h.array[r]) {
		target = r
	}

	if target == idx {
		return
	}

	h.array[target], h.array[idx] = h.array[idx], h.array[target]
	h.heapify(target)
}

// Size returns the size of the heap
//
// NOTE: Size is not indicative of the length of the
// underlying array but rather represents the size of
// the heap portion of the array
func (h *Heap[T]) Size() int {
	return h.size
}
