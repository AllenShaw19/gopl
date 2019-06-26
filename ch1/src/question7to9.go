package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("HTTP Status: %v\n", resp.Status)
		n, err := io.Copy(os.Stdout, resp.Body)
		_ = resp.Body.Close()

		if err != nil || n == 0{
			_, _ = fmt.Fprintf(os.Stderr, "fetch reading %s: %v", url, err)
			os.Exit(1)
		}
	}
}
