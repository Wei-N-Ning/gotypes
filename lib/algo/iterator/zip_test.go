package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZipExpectEmpty(t *testing.T) {
	iter := Zip(Empty[int](), FromSlice([]string{"a", "b"}))
	assert.False(t, iter.Next().IsSome())
}

func TestZipExpectPairs(t *testing.T) {
	var iter Iterator[Pair[int, int]] = Zip(
		Range(0, 10000),
		Range(-100, -95),
	)
	four := iter.Take(4)
	xs := Map(four, func(p Pair[int, int]) int {
		return p.First + p.Second
	}).Slice()
	assert.Equal(t, []int{-100, -98, -96, -94}, xs)
}
