package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func stringTesting() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for r, size := range s {
		fmt.Printf("%c\t%d\n", size, r)
	}
}

func floatTesting() {
	var f float32 = 16777216

	fmt.Printf("%b\n", math.Float32bits(f))
	fmt.Printf("%b\n", math.Float32bits(f+1))
	fmt.Println(f == f+1) // true
}

func main() {

	//stringTesting()
	floatTesting()
}
