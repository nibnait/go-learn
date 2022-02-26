package ch2

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	id   string
	name string
	age  int
}

func (e *Employee) String() string {
	fmt.Printf("实现指针构造方法 Address is %x\n", unsafe.Pointer(&e.name))
	return fmt.Sprintf("ID:%s/name:%s/age:%d", e.id, e.name, e.age)
}

//func (e Employee) String() string {
//	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.name))
//	return fmt.Sprintf("ID:%s-name:%s-age:%d", e.id, e.name, e.age)
//}

func TestCreateEmployeeObj(t *testing.T) {
	e1 := Employee{"0", "Bob", 20}
	t.Log("e1: ", e1)
	t.Log("e1.String: ", e1.String())
	t.Logf("e1 is %T", e1)
	t.Logf("e1.id: %s, e1.name: %s, e1.age: %d", e1.id, e1.name, e1.age)
	println("---")

	e2 := Employee{name: "Mike", age: 30}
	t.Log("e2: ", e2)
	t.Logf("e2 is %T", e2)
	t.Logf("e2.id: %s, e2.name: %s, e2.age: %d", e2.id, e2.name, e2.age)
	println("---")

	e3 := new(Employee) //返回指针
	e3.id = "2"
	e3.age = 22
	e3.name = "Rose"
	t.Log("e3: ", e3)
	t.Logf("e3.id: %s, e3.name: %s, e3.age: %d", e3.id, e3.name, e3.age)
	t.Logf("e3 is %T", e3)
	println("---")

}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.name))
	t.Log(e.String())
}
