package option

import "fmt"

type Option[T any] struct {
	x *T
}

func (opt Option[T]) IsSome() bool {
	return opt.x != nil
}

func (opt Option[T]) ToString() string {
	if opt.IsSome() {
		return fmt.Sprintf("Some(%v)", opt.Unwrap())
	} else {
		return "None"
	}
}

func (opt Option[T]) Unwrap() T {
	return *opt.x
}

// UnwrapOr returns the wrapped value or the default value <defa>
func (opt Option[T]) UnwrapOr(defa T) T {
	if opt.x == nil {
		return defa
	}
	return *opt.x
}

func (opt *Option[T]) Replace(x T) {
	opt.x = &x
}

func Fmap[T any, R any](opt Option[T], f func(x T) R) Option[R] {
	if opt.IsSome() {
		var y R = f(opt.Unwrap())
		return Option[R]{x: &y}
	} else {
		return Option[R]{x: nil}
	}
}

func Some[T any](x T) Option[T] {
	return Option[T]{x: &x}
}

func None[T any]() Option[T] {
	return Option[T]{x: nil}
}
