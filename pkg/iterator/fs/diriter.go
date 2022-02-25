package fs

import (
	"fmt"
	iterator2 "github.com/Wei-N-Ning/gotypes/pkg/iterator"
	. "github.com/Wei-N-Ning/gotypes/pkg/option"
	"io/fs"
	"path/filepath"
)

type Item struct {
	Path     string
	DirEntry fs.DirEntry
}

func DirIter(dir string) iterator2.Iterator[Item] {
	iter, writer := iterator2.TailAppender[Item](1024)
	go func() {
		defer func() {
			writer <- None[Item]()
			close(writer)
		}()
		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			writer <- Some(Item{Path: path, DirEntry: d})
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	}()
	return iter
}