package main

import (
	"fmt"
	"strconv"
	"time"
)

func timeThis(label string, compared int64, f func()) int64 {
	start := time.Now()
	f()
	d := time.Since(start).Milliseconds()
	if compared == 0 {
		fmt.Printf("%32s %16s\n", label, strconv.FormatInt(d, 10))
	} else {
		fmt.Printf("%32s %16s %0.2fx speedup\n", label, strconv.FormatInt(d, 10), float64(compared)/float64(d))
	}
	return d
}
