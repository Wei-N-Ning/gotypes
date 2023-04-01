package vector

// Fill sets every element in the vector to x
func (vec *Vector[T]) Fill(x T) {
	for idx := 0; idx < vec.Size(); idx++ {
		vec.xs[idx] = x
	}
}

// FillRange sets the elements in the given range to x
// If the given range is out of bound, it adjusts it to within [0, size)
// Note the range is exclusive: [from, to)
func (vec *Vector[T]) FillRange(x T, from int, to int) {
	for idx := max(from, 0); idx < min(vec.Size(), to); idx++ {
		vec.xs[idx] = x
	}
}

// FillWith sets every element in the vector to f(x)
func (vec *Vector[T]) FillWith(x T, f func(T) T) {
	for idx := 0; idx < vec.Size(); idx++ {
		vec.xs[idx] = f(x)
	}
}

// FillRangeWith sets the elements in the given range to f(x)
// If the given range is out of bound, it adjusts it to within [0, size)
// Note the range is exclusive: [from, to)
func (vec *Vector[T]) FillRangeWith(x T, from int, to int, f func(T) T) {
	for idx := max(from, 0); idx < min(vec.Size(), to); idx++ {
		vec.xs[idx] = f(x)
	}
}
