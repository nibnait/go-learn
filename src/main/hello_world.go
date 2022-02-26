package main

import (
	"fmt"
	"os"
)

// go run hello_world.go tian bin
func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	//fmt.Println(os.Args[2])

	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}

	fmt.Println("hello world")
	os.Exit(0)
}
