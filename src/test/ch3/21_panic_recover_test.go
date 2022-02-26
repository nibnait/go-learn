package ch3

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestPanicVxExit1(t *testing.T) {
	defer func() {
		fmt.Println("Finally!")
	}()
	fmt.Println("Start")
	//panic(errors.New("Something wrong!"))
	os.Exit(-1)
	fmt.Println("End")
}

func TestPanicVxExit2(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			// 相当于 catch 里面，进行回滚操作
			// ⚠️ 危险！！当⼼ recover 成为恶魔，形成僵⼫服务进程，导致 health check 失效。
			fmt.Println("回滚 ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	//os.Exit(-1)
	fmt.Println("End")
}
