package algo

// to explore the generic fmap(), require go 1.18+

type Output[T any] struct {
	value T
	err   error
}

func GenMap[T any, R any](in <-chan T, f func(x T) R) <-chan Output[R] {
	return nil
}
