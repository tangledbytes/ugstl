package heap

type Store[T any] []T

func (s Store[T]) Parent(idx int) (T, error) {
	var val T

	if idx >= len(s) || idx <= 0 {
		return val, GetErrIndexOutOfBound(idx, len(s))
	}

	return s[parentIdx(idx)], nil
}

func (s Store[T]) Left(idx int) (T, error) {
	var val T

	if idx >= len(s)/2 || idx < 0 {
		return val, GetErrIndexOutOfBound(idx, len(s))
	}

	return s[leftIdx(idx)], nil
}

func (s Store[T]) Right(idx int) (T, error) {
	var val T

	if idx >= len(s)/2 || idx < 0 {
		return val, GetErrIndexOutOfBound(idx, len(s))
	}

	return s[rightIdx(idx)], nil
}

func (s Store[T]) Len() int {
	return len(s)
}

func parentIdx(i int) int {
	return (i - 1) / 2
}

func leftIdx(i int) int {
	return 2*i + 1
}

func rightIdx(i int) int {
	return 2*i + 2
}
