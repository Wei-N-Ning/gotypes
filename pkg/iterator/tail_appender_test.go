package iterator

import (
	"testing"

	. "github.com/Wei-N-Ning/gotypes/pkg/option"
	"github.com/stretchr/testify/assert"
)

func TestTailAppenderExpectGrowth(t *testing.T) {
	iter, writer := TailAppender[int](1024)

	writer <- Some(1)
	writer <- Some(2)
	writer <- None[int]()

	assert.Equal(t, []int{1, 2}, iter.Slice())

	writer <- Some(10)
	writer <- Some(20)
	writer <- None[int]()

	assert.Equal(t, []int{10, 20}, iter.Slice())
}
