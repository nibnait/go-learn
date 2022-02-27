package ch1

import (
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var s string
	t.Logf("s = %s\t\tlen = %d", s, len(s)) //初始化为默认零值“”
	s = "hello"
	t.Logf("s = %s\tlen = %d", s, len(s))

	//s[1] = '3' //string是不可变的byte slice
	//t.Log(s, len(s))

	s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	t.Logf("s = %s\tlen = %d", s, len(s))

	s = "\xE4\xBA\xBB\xFF"
	t.Logf("s = %s\tlen = %d", s, len(s))

}

func TestRune(t *testing.T) {
	// 1. Unicode是一种字符集(codepoint)
	// 2. UTF8是unicode的存储实现（转换为字节序列的规则）

	s := "中"
	t.Logf("s = %s, byte数: %d", s, len(s)) //是byte数
	c := []rune(s)
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

	/*

		字符            "中"
		Unicode     	0x4E2D
		UTF-8       	0xE4B8AD
		string/[]byte	[0xE4,0xB8,0xAD]

	*/
}

func TestStringToRange(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		// [1] 代表第一个参数: c, [2] 代表第二个参数: 22
		t.Logf("%[1]c %[1]x", c)
		t.Logf("%[1]c %[1]d， %[2]d", c, 22)
	}
}
