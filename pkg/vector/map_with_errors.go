package vector

func MapWithErrors[T any, R any](vec Vector[T], f func(T) (R, error)) (Vector[R], Vector[error]) {
	newVec := WithCapacity[R](vec.Capacity())
	errVec := WithCapacity[error](vec.Capacity())
	for idx := 0; idx < vec.size; idx++ {
		if y, err := f(vec.xs[idx]); err == nil {
			newVec.Push(y)
		} else {
			errVec.Push(err)
		}
	}
	return newVec, errVec
}
