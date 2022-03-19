package deque

import (
	vector2 "github.com/Wei-N-Ning/gotypes/pkg/vector"
)

type Deque[T any] struct {
	xs   vector2.Vector[T]
	head int
	tail int
}

func (deq *Deque[T]) Size() int {
	return deq.tail - deq.head + 1
}

func (deq *Deque[T]) Empty() bool {
	return deq.head > deq.tail
}

func WithCapacity[T any](cap uint) Deque[T] {
	return Deque[T]{xs: vector2.WithCapacity[T](int(cap)), head: 0, tail: 0}
}
