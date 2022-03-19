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

func DedupBy[T any](vec Vector[T], cmp func(lhs T, rhs T) bool) Vector[T] {
	if vec.Size() < 2 {
		return vec
	}
	newVec := WithCapacity[T](vec.Capacity() - 1)
	newVec.Push(vec.Head().Unwrap())
	return Fold2[T, Vector[T]](vec.Tail(), newVec, func(v Vector[T], elem T) Vector[T] {
		if !cmp(elem, v.Back()) {
			newVec.Push(elem)
		}
		return newVec
	})
}

func DedupInplaceBy[T any](vec *Vector[T], cmp func(lhs T, rhs T) bool) {
	if vec.Size() < 2 {
		return
	}
	lastIdx := Fold2[T, int](vec.Tail(), 0, func(prevIdx int, elem T) int {
		if cmp(vec.Get(prevIdx), elem) {
			return prevIdx
		}
		prevIdx++
		vec.Set(prevIdx, elem)
		return prevIdx
	})
	vec.size = lastIdx + 1
}

func DedupByKey[T any, K comparable](vec Vector[T], key func(T) K) Vector[T] {
	if vec.Size() < 2 {
		return vec
	}
	newVec := WithCapacity[T](vec.Capacity() - 1)
	newVec.Push(vec.Head().Unwrap())
	return Fold2[T, Vector[T]](vec.Tail(), newVec, func(v Vector[T], elem T) Vector[T] {
		if key(elem) != key(v.Back()) {
			newVec.Push(elem)
		}
		return newVec
	})
}

func DedupInplaceByKey[T any, K comparable](vec *Vector[T], key func(T) K) {
	if vec.Size() < 2 {
		return
	}
	lastIdx := Fold2[T, int](vec.Tail(), 0, func(prevIdx int, elem T) int {
		if key(vec.Get(prevIdx)) == key(elem) {
			return prevIdx
		}
		prevIdx++
		vec.Set(prevIdx, elem)
		return prevIdx
	})
	vec.size = lastIdx + 1
}
