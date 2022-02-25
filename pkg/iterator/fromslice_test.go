package iterator

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFromSliceFilterToSlice(t *testing.T) {
	xs := []string{"abc", "a", "bc", "c", "bbc", ""}
	ys := FromSlice(xs).Filter(func(s string) bool {
		return strings.HasPrefix(s, "a")
	}).Slice()
	ns := FromSlice(xs).Filter(func(s string) bool {
		return strings.HasPrefix(s, "b")
	}).Slice()
	assert.Equal(t, []string{"abc", "a"}, ys)
	assert.Equal(t, []string{"bc", "bbc"}, ns)
}
