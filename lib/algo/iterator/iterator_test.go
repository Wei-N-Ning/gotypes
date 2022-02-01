package iterator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeCount(t *testing.T) {
	iter := Range(1, 100)
	num := iter.Count()
	assert.Equal(t, 100-1, num)
}

func TestRangeFilterCount(t *testing.T) {
	iter := Range(1, 10).Filter(func(x int) bool { return x%2 == 0 })
	num := iter.Count()
	assert.Equal(t, (10-1)/2, num)
}

func TestRangeFilterFold(t *testing.T) {
	iter := Range(1, 10).Filter(func(x int) bool { return x%2 == 0 })
	x := Fold[int, int](iter, 0, func(acc int, elem int) int {
		return acc + elem
	})
	assert.Equal(t, 2+4+6+8, x)
}

func TestRangeMap(t *testing.T) {
	iter := Map(Range(1, 5), func(x int) string {
		return fmt.Sprintf("%02d", x)
	})
	var xs []string
	iter.ForEach(func(x string) {
		xs = append(xs, x)
	})
	assert.Equal(t, []string{"01", "02", "03", "04"}, xs)
}

func TestRepeatTakeCount(t *testing.T) {
	iter := Repeat(0).Take(4)
	x := iter.Count()
	assert.Equal(t, 4, x)
}
