package iterator

import (
	"fmt"
	"strings"
)

func (iter Iterator[T]) String(sep string) string {
	var sb strings.Builder
	Intersperse(
		Map(iter, func(x T) string { return fmt.Sprintf("%v", x) }),
		sep,
	).ForEach(func(s string) {
		sb.WriteString(s)
	})
	return sb.String()
}
