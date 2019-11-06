// 测试包的初始化顺序
package main

import (
	// "ch2/testpkginit"
	"ch2/testpkginit2"
	"fmt"
)

func main() {
	fmt.Printf("%d\n", testpkginit2.A + testpkginit2.M())
}
