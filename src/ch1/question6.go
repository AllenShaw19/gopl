// lissajous 产生随机利萨茹图形的Gif动画
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette__ = []color.Color{color.RGBA{R: 0x00, G: 0x64, B: 0x00, A: 0xff}, color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}}

const (
	GreenIndex__ = 0
	RedIndex__   = 1
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous__(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8008", nil))
		return
	}

	lissajous__(os.Stdout)
}

func lissajous__(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		dalay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette__)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), RedIndex__)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, dalay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
}
