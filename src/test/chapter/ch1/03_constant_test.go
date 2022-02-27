package ch1

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

func TestConstant1(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

// --------------------------------------- //

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant2(t *testing.T) {
	a := 7 // 0111
	t.Log(a&Readable == Readable,
		a&Writable == Writable,
		a&Executable == Executable,
	)
}
