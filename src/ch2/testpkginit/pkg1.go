// 测试包初始化
package testpkginit

import "fmt"

var A = b + c
var b = f()
var c = 1

func f() int{
	return g()
}

func g() int {
	m := 12 + M()
	return m
}

func init()  {
	fmt.Printf("init 1 %d", b)
}

func init() {
	c = 400
}

