package iterator

import "time"

func fib(x int) int {
	if x < 1 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func timeThis(f func()) int64 {
	start := time.Now()
	f()
	return time.Since(start).Nanoseconds()
}
