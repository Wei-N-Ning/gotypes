package vector

import "fmt"

func Window[T any](vec Vector[T], size int) (Vector[*Vector[T]], error) {
	out := WithCapacity[*Vector[T]](0)
	if vec.Empty() {
		return out, fmt.Errorf("vector is empty")
	}
	if size <= 0 || size >= vec.Size() {
		return out, fmt.Errorf("invalid window size: %d (expect a positive integer smaller than the vector size)", size)
	}
	out.Reserve(vec.Size() - size + 1)
	for i := 0; i < vec.Size(); i++ {
		window := WithCapacity[T](size)
		out.Push(&window)
		for j := 0; j < size; j++ {
			if i+j >= vec.Size() {
				return out, nil
			}
			window.Push(vec.xs[i+j])
		}
		if i+size >= vec.Size() {
			break
		}
	}
	return out, nil
}
