package types

// Less is a function signature for the functions which
// must return true when v1 < v2
type Less[T any] func(v1, v2 T) bool
