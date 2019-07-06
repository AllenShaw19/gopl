package main

import (
	"ch2/testpkg"
	"fmt"
)

func main() {
	c := 0.0
	f := testpkg.GetF(c)
	fmt.Println("f:", f)
}
