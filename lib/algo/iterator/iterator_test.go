package iterator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)
import . "go-types-nw/lib/algo/option"

func TestSuccessorFibonacciNumberIterator(t *testing.T) {
	type Item struct {
		first  int
		second int
	}
	init := Some[Item](Item{first: 0, second: 1})
	f := func(x Item) Option[Item] {
		first := x.second
		second := x.first + x.second
		return Some[Item](Item{first: first, second: second})
	}
	iter := Successor(init, f).Take(5)
	var xs []int
	iter.ForEach(func(x Item) {
		xs = append(xs, x.first)
	})
	assert.Equal(t, []int{0, 1, 1, 2, 3}, xs)
}

func TestRangeCount(t *testing.T) {
	iter := Range(1, 100)
	num := iter.Count()
	assert.Equal(t, 100-1, num)
}

func TestRangeFilterCount(t *testing.T) {
	iter := Range(1, 10)
	// the original iterator must be kept;
	// if I do iter = iter.Filter(func(x int) bool { return x%2 == 0 })
	// the original iterator will be (almost) immediately GC-ed which
	// causes its underlying channel to close!!
	// the filter will be blocked forever
	iter_ := iter.Filter(func(x int) bool { return x%2 == 0 })
	num := iter_.Count()
	assert.Equal(t, (10-1)/2, num)
}

func TestRangeFilterFold(t *testing.T) {
	iter := Range(1, 10)
	iter_ := iter.Filter(func(x int) bool { return x%2 == 0 })
	x := Fold[int, int](&iter_, 0, func(acc int, elem int) int {
		return acc + elem
	})
	assert.Equal(t, 2+4+6+8, x)
}

func TestRangeMap(t *testing.T) {
	iter := Range(1, 5)
	iter_ := Map(&iter, func(x int) string {
		return fmt.Sprintf("%02d", x)
	})
	var xs []string
	iter_.ForEach(func(x string) {
		xs = append(xs, x)
	})
	assert.Equal(t, []string{"01", "02", "03", "04"}, xs)
}

func TestRepeatTakeCount(t *testing.T) {
	iter := Repeat(0)
	iter_ := iter.Take(4)
	x := iter_.Count()
	assert.Equal(t, 4, x)
}
