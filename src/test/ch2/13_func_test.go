package ch2

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// ------------- 统计操作时长 ------------------------------ //
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 3)
	return op
}

func TestFn(t *testing.T) {
	//a, b := returnMultiValues()
	//t.Log(a, b)

	time := timeSpent(slowFun)
	t.Log("调用 timeSpent 方法", time(10))
}

// ------------- 可变长度参数 ------------------------------ //

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParm(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
	t.Log(Sum(1, 2, 3, 4))

}

// ------------- defer(延迟)函数（类似于 finally 的作用） ------------------------------ //
func Clear() {
	fmt.Println("3. clear resources 2")
}

func TestDefer(t *testing.T) {
	defer Clear()

	defer func() {
		t.Log("2. clear resources 1")
	}()

	fmt.Println("1. 开始")
	//panic("err")
}
