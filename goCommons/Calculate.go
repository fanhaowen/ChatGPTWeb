package goCommons

import "sync/atomic"

func MyAdd(m *int64) int64 {
	return atomic.AddInt64(m, 1)
}
