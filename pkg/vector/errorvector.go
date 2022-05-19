package vector

import "strings"

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
	xs := Map(errVector, func(err error) string { return err.Error() })
	return AggregatedError{
		inners: errVector,
		sep:    sep,
		joined: strings.Join(xs.ToSlice(), sep),
	}
}
