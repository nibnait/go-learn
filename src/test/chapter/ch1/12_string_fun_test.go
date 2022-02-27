package ch1

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "a,d,c"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}

	t.Logf(strings.Join(parts, "-"))
}

func TestStrConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)

	if atoi, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + atoi)
	}
}
