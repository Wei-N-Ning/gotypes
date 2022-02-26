package vector

func Copy[T any](dst *Vector[T], src Vector[T]) int {
	return CopyAt(dst, src, 0)
}

func CopyAt[T any](dst *Vector[T], src Vector[T], at int) int {
	limit := dst.Size()
	copied := 0
	for idx := 0; idx < src.Size(); idx++ {
		if at < limit {
			dst.xs[at] = src.xs[idx]
			at += 1
			copied += 1
		} else {
			break
		}
	}
	return copied
}
