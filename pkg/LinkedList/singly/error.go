package singly

import "errors"

var (
	ErrIndexOutOfBound = errors.New("index is out of bound")
	ErrInvalidState    = errors.New("invalid state")
)
