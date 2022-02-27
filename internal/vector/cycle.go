package vector

func Cycle[T any](vec Vector[T], num int) Vector[T] {
	if num <= 0 || vec.Empty() {
		return vec
	}
	return Concat[T](vec, Repeat[Vector[T]](vec, num-1).ToSlice()...)
}
