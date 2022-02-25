package iterator

type Pair[T, P any] struct {
	First  T
	Second P
}

func (p Pair[T, P]) Unpack() (T, P) {
	return p.First, p.Second
}

func NewPair[T, P any](x T, y P) Pair[T, P] {
	return Pair[T, P]{First: x, Second: y}
}
