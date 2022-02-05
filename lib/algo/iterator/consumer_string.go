package iterator

import (
	"bytes"
)

func String[T any](iter Iterator[T], sep string) string {
	var buffer bytes.Buffer

	return buffer.String()
}
