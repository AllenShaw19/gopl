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
	"strconv"
	"time"
)

var palette2 = []color.Color{color.White, color.Black}

const (
	whiteIndex2 = 0
	blackIndex2 = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			cycles := 5
			if len(r.Form["cycles"]) > 0 {
				inCycles, err := strconv.Atoi(r.Form["cycles"][0])
				if err != nil {
					fmt.Println(err)
					inCycles = 5
				}
				cycles = inCycles
			}

			lissajous12(w, float64(cycles))
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8008", nil))
		return
	}

	lissajous12(os.Stdout, 5.0)
}

func lissajous12(out io.Writer, cycles float64) {
	const (
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
		img := image.NewPaletted(rect, palette2)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex2)
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
