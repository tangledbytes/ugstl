package utils

import "testing"

func Reverse[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func ShouldPanic(t *testing.T, f func()) {
	t.Helper()
	defer func() { recover() }()
	f()

	t.Error("function must have panicked")
}
