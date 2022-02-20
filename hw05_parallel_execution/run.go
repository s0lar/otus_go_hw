package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	taskCh := make(chan Task)
	var counter int32

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				taskResult, ok := <-taskCh
				if !ok {
					break
				}
				err := taskResult()
				if err != nil {
					atomic.AddInt32(&counter, 1)
				}
			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&counter) == int32(m) {
			break
		}
		taskCh <- task
	}

	close(taskCh)

	wg.Wait()

	if atomic.LoadInt32(&counter) >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
