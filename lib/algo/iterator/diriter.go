package iterator

import (
	"fmt"
	. "go-types-nw/lib/algo/option"
	"io/fs"
	"path/filepath"
)

type Item struct {
	Path     string
	DirEntry fs.DirEntry
}

func dirIterImpl(dir string) <-chan Option[Item] {
	ch := make(chan Option[Item])
	go func() {
		defer func() {
			ch <- None[Item]()
			close(ch)
		}()
		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			ch <- Some(Item{Path: path, DirEntry: d})
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	}()
	return ch
}

func DirIter(dir string) Iterator[Item] {
	return Iterator[Item]{ch: dirIterImpl(dir)}
}
