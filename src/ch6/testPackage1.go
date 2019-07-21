package main

import "ch6/testch"
import "fmt"

func main() {
	testch.Test()

	p := testch.Package{
		1, 3,
	}

	fmt.Println(p)
	//p.x = 15
}
