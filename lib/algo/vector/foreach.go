package vector

func (vec *Vector[T]) ForEach(f func(T)) {
	for idx := 0; idx < vec.size; idx++ {
		f(vec.xs[idx])
	}
}
