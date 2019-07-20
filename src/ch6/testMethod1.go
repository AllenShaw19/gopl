package main

import "fmt"

type Int int

func (i Int) f() {
	fmt.Println(i)
}

func main() {
	var i Int
	i = 10
	i.f()
}
