package vector

func Fold[T any](vec Vector[T], init T, f func(acc T, elem T) T) T {
	vec.ForEach(func(x T) {
		init = f(init, x)
	})
	return init
}
