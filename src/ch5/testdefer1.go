package main

import "fmt"

func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	x = x + x
	return x
}

func main() {
	double(3)
}
