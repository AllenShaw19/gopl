package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	i := reflect.ValueOf(3)

	inf := i.Interface()

	fmt.Printf("%T, %v\n", t, t)
	fmt.Printf("%T, %v\n", i, i)
	fmt.Printf("%T, %v\n", inf, inf)
}
