package ch4

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask1(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)

	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask1(i)
			ch <- ret
		}(i)
	}

	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
