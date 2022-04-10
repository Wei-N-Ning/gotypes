package vector

func Repeat[T any](init T, num int) Vector[T] {
	out := WithCapacity[T](max(0, num))
	for idx := 0; idx < num; idx++ {
		out.Push(init)
	}
	return out
}
