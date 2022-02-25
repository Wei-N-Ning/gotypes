package vector

func Map[T any, R any](vec Vector[T], f func(T) R) Vector[R] {
	newVec := WithCapacity[R](vec.Capacity())
	for idx := 0; idx < vec.size; idx++ {
		newVec.xs[idx] = f(vec.xs[idx])
	}
	newVec.size = vec.size
	return newVec
}
