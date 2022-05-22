package vector

import (
	"strings"

	"github.com/Wei-N-Ning/gotypes/pkg/option"
)

type ErrVector = Vector[error]

type AggregatedError struct {
	inners ErrVector
	sep    string
	joined string
}

func (e AggregatedError) Error() string {
	return e.joined
}

func AggregateError(errVector ErrVector, sep string) error {
	if errVector.Empty() {
		return nil
	}
	xs := MapFilter(errVector, func(err error) option.Option[string] {
		if err == nil {
			return option.None[string]()
		}
		return option.Some[string](err.Error())
	})
	if xs.Empty() {
		return nil
	}
	return AggregatedError{
		inners: errVector,
		sep:    sep,
		joined: strings.Join(xs.ToSlice(), sep),
	}
}
