package ch2

import "testing"

type Programmer1 interface {
	writeHelloWorld() string
}

type GoProgrammer1 struct {
}

func (g *GoProgrammer1) writeHelloWorld() string {
	return "fmt.Println(\"hello world\")"
}

func TestClient(t *testing.T) {
	var p Programmer1
	p = new(GoProgrammer1)
	t.Log(p.writeHelloWorld())
}
