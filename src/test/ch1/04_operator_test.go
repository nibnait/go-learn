package ch1

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 3}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	t.Log(a == d)

	// 长度不同的数据，无法比较
	// t.Log(a == c)

}

// 按位清零
func TestBitClear(t *testing.T) {
	a := 7 // 0111
	t.Log(a&Readable == Readable)

	a = a &^ Readable
	t.Log(a&Readable == Readable)
}
