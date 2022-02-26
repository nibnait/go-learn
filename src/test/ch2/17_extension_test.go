package ch2

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) speak() {
	fmt.Print("...")
}

func (p *Pet) speakTo(host string) {
	p.speak()
	fmt.Println("  ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) speak() {
	fmt.Println("汪！汪！")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.speakTo("tian")
}
