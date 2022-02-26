package ch2

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	writeHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) writeHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) writeHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("p.type: %T\nret: %v\n\n", p, p.writeHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
