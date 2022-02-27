package deque

import "github.com/Wei-N-Ning/gotypes/internal/vector"

type Deque[T any] struct {
	xs   vector.Vector[T]
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
	return Deque[T]{xs: vector.WithCapacity[T](int(cap)), head: 0, tail: 0}
}
