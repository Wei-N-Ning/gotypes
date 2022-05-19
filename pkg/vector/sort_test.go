package vector

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSortVector(t *testing.T) {
	t.Run("empty vector, expect no effect", func(t *testing.T) {
		v := FromValues[int]()
		Sort(&v)
		assert.True(t, v.Empty())
		v1 := FromValues(1)
		Sort(&v1)
		assert.Equal(t, []int{1}, v1.ToSlice())
	})
	t.Run("expect element order", func(t *testing.T) {
		v := FromValues(3, 1, 4, 1, 5, 9, 2, 6)
		Sort(&v)
		assert.Equal(t, []int{1, 1, 2, 3, 4, 5, 6, 9}, v.ToSlice())
	})
}

func TestVectorSortBy(t *testing.T) {
	t.Run("expect custom less order", func(t *testing.T) {
		v := FromValues("1", "x", "2a", "3", "aaa", "4")
		SortBy(&v, func(i, j int) bool {
			lhs, lErr := strconv.ParseInt(v.Get(i), 10, 32)
			rhs, rErr := strconv.ParseInt(v.Get(j), 10, 32)
			if lErr == nil {
				if rErr == nil {
					return lhs < rhs
				} else {
					return false
				}
			} else {
				if rErr == nil {
					return true
				} else {
					return lhs < rhs
				}
			}
		})
		fmt.Println(v.ToSlice())
	})
}
