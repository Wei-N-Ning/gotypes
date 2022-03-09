package vector

// Dedup removes the duplicate contiguous elements, e.g.
// [1, 2, 2, 2, 3, 3, 4]
// dedup to:
// [1, 2, 3, 4]
// but for [1, 2, 3, 2, 3, 2]
// dedup has no effect
func Dedup[T comparable](vec Vector[T]) Vector[T] {
	if vec.Size() < 2 {
		return vec
	}
	newVec := WithCapacity[T](vec.Capacity() - 1)
	newVec.Push(vec.Head().Unwrap())
	return Fold2[T, Vector[T]](vec.Tail(), newVec, func(v Vector[T], elem T) Vector[T] {
		if elem != v.Back() {
			newVec.Push(elem)
		}
		return newVec
	})
}

// DedupInplace will modify the given vector
func DedupInplace[T comparable](vec *Vector[T]) {
	if vec.Size() < 2 {
		return
	}
	lastIdx := Fold2[T, int](vec.Tail(), 0, func(prevIdx int, elem T) int {
		if vec.Get(prevIdx) == elem {
			return prevIdx
		}
		prevIdx++
		vec.Set(prevIdx, elem)
		return prevIdx
	})
	vec.size = lastIdx + 1
}
