package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEachWithIndex(t *testing.T) {
	iter := FromSlice([]string{"a", "b", "c"})
	var xs []int
	iter.ForEachWithIndex(func(idx int, elem string) {
		xs = append(xs, idx)
	})
	assert.Equal(t, []int{0, 1, 2}, xs)
}
