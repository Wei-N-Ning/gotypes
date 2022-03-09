package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDedup(t *testing.T) {
	t.Run("empty vector", func(t *testing.T) {
		vs := FromValues[int]()
		dvs := Dedup(vs)
		assert.Equal(t, []int{}, dvs.ToSlice())
	})
	t.Run("expect elements removed", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3}, Dedup(FromValues(1, 2, 2, 3)).ToSlice())
		assert.Equal(t, []int{1}, Dedup(FromValues(1, 1, 1, 1)).ToSlice())
	})
	t.Run("expect no effect", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 4, 2, 3}, Dedup(FromValues(1, 2, 4, 2, 3)).ToSlice())
	})
}

func TestDedupBy(t *testing.T) {
	cmp := func(lhs string, rhs string) bool {
		return lhs[:min(2, len(lhs))] == rhs[:min(2, len(rhs))]
	}
	t.Run("empty vector", func(t *testing.T) {
		vs := FromValues[string]()
		dvs := DedupBy(vs, cmp)
		assert.Equal(t, []string{}, dvs.ToSlice())
	})
	t.Run("expect elements removed", func(t *testing.T) {
		assert.Equal(t, []string{"a", "idkfa", "b"}, DedupBy(FromValues("a", "idkfa", "idfa", "b"), cmp).ToSlice())
		assert.Equal(t, []string{"id"}, DedupBy(FromValues("id", "idnoclip", "iddqd", "idkfa", "idfa"), cmp).ToSlice())
	})
	t.Run("expect no effect", func(t *testing.T) {
		assert.Equal(t, []string{"a", "b", "c"}, DedupBy(FromValues("a", "b", "c"), cmp).ToSlice())
	})
}

func TestDedupByKey(t *testing.T) {
	keyF := func(x string) int {
		return len(x)
	}
	t.Run("empty vector", func(t *testing.T) {
		vs := FromValues[string]()
		dvs := DedupByKey(vs, keyF)
		assert.Equal(t, []string{}, dvs.ToSlice())
	})
	t.Run("expect elements removed", func(t *testing.T) {
		assert.Equal(t, []string{"1", "e1m1", "b"}, DedupByKey(FromValues("1", "a", "e1m1", "idfa", "b", "_"), keyF).ToSlice())
		assert.Equal(t, []string{"id"}, DedupByKey(FromValues("id", "e1", "m1", "zz", "op"), keyF).ToSlice())
	})
	t.Run("expect no effect", func(t *testing.T) {
		assert.Equal(t, []string{"a1", "b12", "c"}, DedupByKey(FromValues("a1", "b12", "c"), keyF).ToSlice())
	})
}

func TestDedupInplace(t *testing.T) {
	t.Run("empty vector", func(t *testing.T) {
		vs := FromValues[int]()
		DedupInplace(&vs)
		assert.Equal(t, []int{}, vs.ToSlice())
	})
	t.Run("expect elements removed", func(t *testing.T) {
		vs := FromValues[int](1, 1, 2, 2, 3, 4, 4, 5)
		DedupInplace(&vs)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, vs.ToSlice())
	})
	t.Run("expect elements reduced to singleton", func(t *testing.T) {
		vs := FromValues[int](1, 1, 1, 1, 1, 1, 1, 1)
		DedupInplace(&vs)
		assert.Equal(t, []int{1}, vs.ToSlice())
	})
}

func TestDedupInplaceBy(t *testing.T) {
	cmp := func(lhs string, rhs string) bool {
		return lhs[:min(2, len(lhs))] == rhs[:min(2, len(rhs))]
	}
	t.Run("empty vector", func(t *testing.T) {
		vs := FromValues[string]()
		DedupInplaceBy(&vs, cmp)
		assert.Equal(t, []string{}, vs.ToSlice())
	})
	t.Run("expect elements removed", func(t *testing.T) {
		vs := FromValues("aab", "aa", "idkfa", "idfa", "bl", "bliz")
		DedupInplaceBy(&vs, cmp)
		assert.Equal(t, []string{"aab", "idkfa", "bl"}, vs.ToSlice())
	})
	t.Run("expect elements reduced to singleton", func(t *testing.T) {
		vs := FromValues[int](1, 1, 1, 1, 1, 1, 1, 1)
		DedupInplaceBy(&vs, func(x int, y int) bool { return x == y })
		assert.Equal(t, []int{1}, vs.ToSlice())
	})
}

func TestDedupInplaceByKey(t *testing.T) {
	t.Run("empty vector", func(t *testing.T) {
		vs := FromValues[string]()
		DedupInplaceByKey(&vs, func(string) int { return 0 })
		assert.Equal(t, []string{}, vs.ToSlice())
	})
	t.Run("expect elements removed", func(t *testing.T) {
		vs := FromValues("a_", "aa", "e1m1", "idfa", "e1m4", "bliz")
		DedupInplaceByKey(&vs, func(x string) int { return len(x) })
		assert.Equal(t, []string{"a_", "e1m1"}, vs.ToSlice())
	})
	t.Run("expect elements reduced to singleton", func(t *testing.T) {
		vs := FromValues[int](1, 1, 1, 1, 1, 1, 1, 1)
		DedupInplaceByKey(&vs, func(x int) bool { return x == 10 })
		assert.Equal(t, []int{1}, vs.ToSlice())
	})
}
