package vector

func (vec *Vector[T]) Extend(other Vector[T]) {
	vec.Reserve(other.Size())
	other.ForEach(func(elem T) {
		vec.Push(elem)
	})
}

func (vec *Vector[T]) ExtendBy(elem T, num int) {
	vec.Reserve(num)
	for idx := 0; idx < num; idx++ {
		vec.Push(elem)
	}
}
