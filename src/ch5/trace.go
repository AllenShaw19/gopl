package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()()

	fmt.Println("Sleep 3 Second Before")
	time.Sleep(3 * time.Second)
	fmt.Println("Sleep 3 Second After")
}

func trace(msg string) func() func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
		return func() {
			log.Printf("hello world")
		}
	}
}

func main() {
	bigSlowOperation()
}
