package testch

import "fmt"

func Test() {
	p := &Package{
		x: 1,
		y: 2,
	}

	fmt.Println(p.x)

	p.x = 133
	fmt.Println(p.x)
}
