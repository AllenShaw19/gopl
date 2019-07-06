package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	path := "./urlcontent.txt"
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch2File(url, ch, f)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	_ = f.Close()

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch2File(url string, ch chan string, w io.Writer) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	n, err := io.Copy(w, resp.Body)
	_ = resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s, %v", url, err)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, n, url)
}
