package vector

import "fmt"

// chunks:
// given 1, 2, 3, ... n, chunk_size = 3
// return [1, 2, 3], [4, 5, 6], ... [n-2, n-1, n]

// window (aka sliding window):
// given 1, 2, 3, ... n, window_size = 3
// return [1, 2, 3], [2, 3, 4], ... [n-3, n-2, n-1], [n-2, n-1, n]

func Chunks[T any](vec Vector[T], size int) (Vector[*Vector[T]], error) {
	out := WithCapacity[*Vector[T]](0)
	if size <= 0 {
		return out, fmt.Errorf("invalid chunk size: %d (expect a positive integer)", size)
	}
	if vec.Empty() {
		return out, fmt.Errorf("vector is empty")
	}
	out.Reserve(vec.Size()/size + 1)
	chunk := WithCapacity[T](size)
	out.Push(&chunk)
	vec.ForEach(func(elem T) {
		if chunk.Size() == size {
			chunk = WithCapacity[T](size)
			out.Push(&chunk)
		}
		chunk.Push(elem)
	})
	return out, nil
}
