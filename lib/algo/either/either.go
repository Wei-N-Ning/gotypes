package either

type Either[T any] struct {
	x   *T
	err error
}

func (et *Either[T]) IsOk() bool {
	return et.err == nil
}

func (et Either[T]) Unwrap() T {
	return *et.x
}

func (et Either[T]) UnwrapErr() error {
	return et.err
}

func Ok[T any](x T) Either[T] {
	return Either[T]{x: &x, err: nil}
}

func Err[T any](err error) Either[T] {
	return Either[T]{x: nil, err: err}
}

func Fmap[T any, R any](et Either[T], f func(x T) R) Either[R] {
	if et.IsOk() {
		y := f(et.Unwrap())
		return Either[R]{x: &y, err: nil}
	} else {
		return Either[R]{x: nil, err: et.UnwrapErr()}
	}
}
