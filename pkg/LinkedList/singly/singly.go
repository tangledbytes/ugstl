package singly

import (
	iter "github.com/utkarsh-pro/ugstl/pkg/Iter"
)

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  uint
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) IsEmpty() bool {
	return l.head == nil
}

func (l *List[T]) Len() uint {
	return l.len
}

func (l *List[T]) InsertAtHead(val T) {
	node := NewNode(val)
	l.len++

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.SetNext(l.head)
	l.head = node
}

func (l *List[T]) InsertAtTail(val T) {
	node := NewNode(val)
	l.len++

	if l.tail == nil {
		l.head = node
		l.tail = node
		return
	}

	l.tail.SetNext(node)
	l.tail = node
}

func (l *List[T]) InsertAtN(val T, n uint) error {
	if n >= l.Len() {
		return ErrIndexOutOfBound
	}

	if n == 0 {
		l.InsertAtHead(val)
		return nil
	}

	if n == l.Len()-1 {
		l.InsertAtTail(val)
		return nil
	}

	node := NewNode(val)

	prevNode, ok := l.Get(n - 1)
	if !ok {
		return ErrInvalidState // This case should never happen
	}

	prevNodeNext := prevNode.GetNext()
	prevNode.SetNext(node)
	node.SetNext(prevNodeNext)

	return nil
}

func (l *List[T]) Iter() iter.Iter[*Node[T]] {
	return NewIter(l)
}

func (l *List[T]) Get(n uint) (*Node[T], bool) {
	if n >= l.Len() {
		return nil, false
	}

	var node *Node[T]

	iter := l.Iter()
	for i := uint(0); i <= n; i++ {
		node, _ = iter.Next()
	}

	return node, true
}

func (l *List[T]) RemoveFromHead() {
	if l.head == nil {
		return
	}

	l.len--
	l.head = l.head.GetNext()
}

func (l *List[T]) RemoveAtN(n uint) error {
	if n >= l.Len() {
		return ErrIndexOutOfBound
	}

	if n == 0 {
		l.RemoveFromHead()
		return nil
	}

	prevNode, ok := l.Get(n - 1)
	if !ok {
		return ErrInvalidState
	}

	node := prevNode.GetNext()
	if node == nil {
		return ErrInvalidState
	}

	l.len--
	prevNode.SetNext(node.GetNext())

	return nil
}
