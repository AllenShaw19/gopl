package testpkginit2

import "fmt"

var A int
var a = b + c
var b = f()
var c = 1
var d = 2
func f() int { return c + 1 }

func init() {
	fmt.Println("init A = a")
	A = a
}

func init() {
	fmt.Println("init d = ")
	d = c
}

func M() int {
	return b
}
