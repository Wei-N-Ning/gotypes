package vector

import (
	"github.com/Wei-N-Ning/gotypes/pkg/option"
	"github.com/stretchr/testify/assert"
	"path"
	"strings"
	"testing"
)

func TestExpectElementsFiltered(t *testing.T) {
	vec := FromValues(
		"/tmp",
		"/var/tmp",
		"/home/user",
		"home/user/tmp",
	)
	xs := MapFilter(vec, func(s string) option.Option[string] {
		if strings.HasSuffix(s, "tmp") && strings.HasPrefix(s, "/") {
			return option.Some(path.Dir(s))
		} else {
			return option.None[string]()
		}
	})
	assert.Equal(t, []string{"/", "/var"}, xs.ToSlice())
}