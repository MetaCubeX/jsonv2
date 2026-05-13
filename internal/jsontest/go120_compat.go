package jsontest

import "sync"

func onceValue[T any](f func() T) func() T {
	var (
		once sync.Once
		v    T
	)
	return func() T {
		once.Do(func() {
			v = f()
		})
		return v
	}
}
