package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func testPointer() {
	var p = f()
	fmt.Println("p", p, *p)
}

func main() {
	testPointer()
}
