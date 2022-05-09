package heap

import "github.com/utkarsh-pro/ugstl/pkg/types"

// Sort will sort the given array either in ascending order
func Sort[T any](arr []T, less types.Less[T]) {
	h := NewFromArray(arr, MAX, less)

	for h.size > 0 {
		h.array[0], h.array[h.size-1] = h.array[h.size-1], h.array[0]
		h.size--
		h.heapify(0)
	}
}
