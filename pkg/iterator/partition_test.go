package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionExpectBothEmpty(t *testing.T) {
	l, r := Partitioned(Empty[int](), func(x int) bool { return x%2 == 1 })
	assert.False(t, l.Next().IsSome())
	assert.False(t, r.Next().IsSome())
}

func TestPartitionExpectOrder(t *testing.T) {
	l, r := Partitioned(Range(1, 100), func(x int) bool { return x%2 == 1 })
	xs := l.Take(4).Slice()
	ys := r.Take(4).Slice()
	assert.Equal(t, []int{1, 3, 5, 7}, xs)
	assert.Equal(t, []int{2, 4, 6, 8}, ys)
}
