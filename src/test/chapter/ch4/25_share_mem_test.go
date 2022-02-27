package ch4

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {

	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)

}

// lock unlock --> sync.Mutex
func TestCounterThreadSafeMutex(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()

			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// lock unlock --> sync.RWLock ✅ 读多写少时，性能更好
func TestCounterThreadSafeRWLock(t *testing.T) {
	var rwMutex sync.RWMutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				rwMutex.Unlock()
			}()

			rwMutex.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// thread.join -> sync.WaitGroup(.Add(1), .Done, .Wait)
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()

			mut.Lock()
			counter++

			wg.Done()
		}()
	}

	wg.Wait()
	t.Logf("counter = %d", counter)
}
