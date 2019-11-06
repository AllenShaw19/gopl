// 测试包初始化顺序
package testpkginit

import "fmt"

var d = A + 1

func M() int {
	return d
}

func init()  {
	fmt.Printf("%d", M())
}