package main

import (
	"fmt"
	"os"
	"runtime"
)

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	PrintStack()
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], true)
	_, _ = os.Stdout.Write(buf[:n])
}

func main() {
	defer PrintStack()
	f(3)
}
