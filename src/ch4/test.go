// map 与 slice 的一些不同
package main

import "fmt"

func main() {
	var a []int //TODO: 创建一个nil的slice，底层没有对应的数组
	if a == nil {
		fmt.Println("No.1 a == nil", a)
	}

	//TODO: nil 的slice可以append，但是不能获取a[2]，slice获取超过底层数组长度的值时会失败
	a = append(a, 123)

	if a == nil {
		fmt.Println("No.2 a == nil", a)
	}

	var b map[string]int //TODO 创建一个nil的map，底层没有对应的哈希表
	b = map[string]int{}
	//TODO 与slice不同，nil的map可以获取a["a"],会返回零值，但是不可以插入新数据，a["a"]=123则会失败
	b["a"] = 1

}
