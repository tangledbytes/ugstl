package sort

import (
	"github.com/utkarsh-pro/ugstl/pkg/heap"
	"github.com/utkarsh-pro/ugstl/pkg/types"
)

func Heap[T any](arr []T, less types.Less[T]) {
	heap.Sort(arr, less)
}

func Selection[T any](arr []T, less types.Less[T]) {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i; j < len(arr); j++ {
			if less(arr[j], arr[minIdx]) {
				minIdx = j
			}
		}

		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}
}
