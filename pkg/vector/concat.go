package vector

func Concat[T any](head Vector[T], tails ...Vector[T]) Vector[T] {
	newCap := head.Capacity()
	for idx := range tails {
		newCap += tails[idx].Capacity()
	}
	head.reallocate(newCap)
	for idx := range tails {
		head.Extend(tails[idx])
	}
	return head
}
