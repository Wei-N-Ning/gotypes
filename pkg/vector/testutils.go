package vector

import "time"

func timeThis(f func()) int64 {
	start := time.Now()
	f()
	return time.Since(start).Nanoseconds()
}
