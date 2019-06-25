package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, path := range files {
			f, err := os.Open(path)
			if err != nil {
				_ = fmt.Errorf("dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			_ = f.Close()
		}
	}
	printCounts(counts)
}

func printCounts(counts map[string]int) {
	for k, v := range counts {
		if v > 0 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		counts[input.Text()]++
	}
}
