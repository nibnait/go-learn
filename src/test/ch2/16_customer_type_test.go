package ch2

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpentV2(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func TestFnV2(t *testing.T) {
	tsSF := timeSpentV2(slowFun)
	t.Log(tsSF(10))
}
