package iterator

func (iter Iterator[T]) Slice() []T {
	var xs []T
	iter.ForEach(func(x T) {
		xs = append(xs, x)
	})
	return xs
}
