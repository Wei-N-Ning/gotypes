package iterator

func (iter Iterator[T]) Count() int {
	num := 0
	for {
		elem := iter.Next()
		if elem.IsSome() {
			num += 1
		} else {
			break
		}
	}
	return num
}
