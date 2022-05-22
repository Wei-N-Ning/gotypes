package iterator

import (
	"testing"

	. "github.com/Wei-N-Ning/gotypes/pkg/option"
	"github.com/stretchr/testify/assert"
)

func TestFibonacciNumberSeries(t *testing.T) {
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
