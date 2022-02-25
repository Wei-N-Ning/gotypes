package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenEmptyIterators(t *testing.T) {
	iterators := Map(Repeat("a").Take(5), func(string) Iterator[int] {
		return Empty[int]()
	})
	iter := Flatten(iterators)
	assert.False(t, iter.Next().IsSome())
}

func TestFlattenIteratorWithValues(t *testing.T) {
	iterators := Map(
		FromSlice([]string{"a", "12", "___"}),
		func(s string) Iterator[string] {
			return RepeatN("b", len(s))
		},
	)
	xs := Flatten(iterators).Slice()
	assert.Equal(t, []string{"b", "b", "b", "b", "b", "b"}, xs)
}
