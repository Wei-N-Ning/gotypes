package iterator

// Chain is a simpler form of Concat:
// It concatenates itself and another iterator of the same type.
func (iter Iterator[T]) Chain(other Iterator[T]) Iterator[T] {
	return Concat(iter, other)
}
