package syncbits

import (
	"sync"
)

// Workgroup creates a WaitGroup that executes the given function with
// the specified amount of workers and completes once all the workers are done.
func Workgroup(fn func(), workers int) *sync.WaitGroup {
	wg := sync.WaitGroup{}

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			fn()
		}()
		wg.Add(1)
	}

	return &wg
}
