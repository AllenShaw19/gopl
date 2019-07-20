package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		surface(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8008", nil))
	return
}

func surface(out io.Writer) {

}
