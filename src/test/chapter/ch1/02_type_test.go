package ch1

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	// b = a  ❌ 不支持
	b = int64(a)

	t.Log(a, b)

	var c MyInt
	// c = b  ❌ 不支持
	c = MyInt(b)
	t.Log(c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T, %T", a, aPtr)

	// 不支持指针运算
	// aPtr = aPtr + 1
}

func TestStringInit(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))

	//if s == nil {
	//	t.Log("s == nil")
	//}

	if s == "" {
		t.Log("s == \"\"")
	}

}
