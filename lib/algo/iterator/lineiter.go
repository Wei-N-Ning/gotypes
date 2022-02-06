package iterator

import (
	"bufio"
	"fmt"
	. "go-types-nw/lib/algo/option"
	"io"
)

type Line struct {
	Num   int
	Value string
}

func (l *Line) ToString() string {
	return fmt.Sprintf("%d: %s", l.Num, l.Value)
}

func Lines(reader io.Reader) Iterator[Line] {
	iter, writer := TailAppender[Line](1024)
	go func() {
		defer func() {
			writer <- None[Line]()
			close(writer)
		}()
		scanner := bufio.NewScanner(reader)
		lineNum := 1
		for scanner.Scan() {
			writer <- Some(Line{Num: lineNum, Value: scanner.Text()})
			lineNum += 1
		}
	}()
	return iter
}
