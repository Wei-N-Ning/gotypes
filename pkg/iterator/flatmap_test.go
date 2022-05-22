package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatMapExpectElements(t *testing.T) {
	xs := FlatMap(
		RangeInclusive(1, 3),
		func(x int) Iterator[string] {
			return RepeatN("a", x)
		},
	).Slice()
	assert.Equal(t, []string{"a", "a", "a", "a", "a", "a"}, xs)
}
