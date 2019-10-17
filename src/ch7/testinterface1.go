package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	var w io.Writer
	var j interface{}
	w = os.Stdout
	m := j.(interface{})
	d := w.(interface{})
	f := w.(*os.File)
	_, _ = w.(*bytes.Buffer)
	println("---------------", m, d, f)
}
