package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func the_for_method() {
	s, sep := "", ""
	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("the_for_method %.2fs elapsed\n", time.Since(start).Nanoseconds())
}

func the_join_method() {
	start := time.Now()
	s := strings.Join(os.Args[1:], " ")
	fmt.Println(s)

	fmt.Printf("the_join_method %.2fs elapsed\n", time.Since(start).Nanoseconds())
}

func main() {

	the_for_method()

	the_join_method()

}
