package iterator

func All[T any](iter Iterator[T], f func(x T) bool) bool {
	out := true
	for {
		elem := iter.Next()
		if elem.IsSome() {
			out = out && f(elem.Unwrap())
			if !out {
				return false
			}
		} else {
			break
		}
	}
	return out
}

func Any[T any](iter Iterator[T], f func(x T) bool) bool {
	out := true
	for {
		elem := iter.Next()
		if elem.IsSome() {
			out = out || f(elem.Unwrap())
			if out {
				return true
			}
		} else {
			break
		}
	}
	return out
}
