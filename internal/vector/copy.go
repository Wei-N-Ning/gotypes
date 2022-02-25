package vector

func Copy[T any](dst *Vector[T], src Vector[T]) int {
	limit := dst.Size()
	copied := 0
	for idx, elem := range src.xs {
		if idx < limit {
			dst.xs[idx] = elem
			copied += 1
		} else {
			break
		}
	}
	return copied
}

func (vec *Vector[T]) CopyFrom(other Vector[T]) int {
	return Copy(vec, other)
}
