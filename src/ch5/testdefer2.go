package main

import (
	"fmt"
	"time"
)

func main() {

	var s []int
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}

	for j := range s {
		defer func() {
			fmt.Printf("defer func %d\n", j)
		}()

		fmt.Printf("do sometime %d\n", j)
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("--------------- Main End!!! ---------------")

}
