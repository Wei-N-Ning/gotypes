package iterator

func Empty[T any]() Iterator[T] {
	return FromSlice[T]([]T{})
}
