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
