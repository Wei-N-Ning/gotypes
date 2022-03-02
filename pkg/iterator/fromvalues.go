package iterator

func FromValues[T any](xs ...T) Iterator[T] {
	return FromSlice(xs)
}
