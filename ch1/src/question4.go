package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Record struct {
	Count     int
	FileNames map[string]bool
}

func main() {
	dupFiles := make(map[string]Record)
	files := os.Args[1:]

	if len(files) == 0 {
		record(os.Stdin, dupFiles)
	} else {
		for _, path := range files {
			f, err := os.Open(path)
			if err != nil {
				_ = fmt.Errorf("question4 %v", err)
			}
			record(f, dupFiles)
			_ = f.Close()
		}
	}

	printDupFiles(dupFiles)
}

func printDupFiles(dupFiles map[string]Record) {
	for k, v := range dupFiles {
		if v.Count > 0 {
			var fileNames []string
			for fileName, _ := range v.FileNames {
				fileNames = append(fileNames, fileName)
			}
			fmt.Printf("%v\t%d\t%s\n", strings.Join(fileNames, " "), v.Count, k)
		}
	}
}

func record(file *os.File, dupFiles map[string]Record) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		if input.Text() == "" {
			break
		}
		if v, ok := dupFiles[input.Text()]; ok {
			v.Count++
			v.FileNames[file.Name()] = true
		} else {
			dupFiles[input.Text()] = Record{Count: 1, FileNames: map[string]bool{file.Name(): true}}
		}
	}

}
