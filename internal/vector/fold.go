package vector

func Fold[T any](vec Vector[T], init T, f func(acc T, elem T) T) T {
	vec.ForEach(func(x T) {
		init = f(init, x)
	})
	return init
}

func Fold2[T, P any](vec Vector[T], init P, f func(acc P, elem T) P) P {
	vec.ForEach(func(x T) {
		init = f(init, x)
	})
	return init
}
