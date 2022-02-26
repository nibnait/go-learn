package ch2

import (
	"fmt"
	"testing"
)

func doSth(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("int", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("string", i)
		return
	}
	fmt.Println("unknown type")
}

func doSthSwitch(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknown type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	doSthSwitch(10)
	doSthSwitch("xxx")
}
