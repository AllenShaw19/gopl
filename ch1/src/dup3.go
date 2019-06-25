package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	for _, path := range files {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			_ = fmt.Errorf("dup3 %v\n", err)
			continue
		}

		for _ ,v := range strings.Split(string(data), "\n") {
			counts[v]++
		}
	}

	for k, v := range counts {
		if v > 0 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}
