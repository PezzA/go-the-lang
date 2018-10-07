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
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.White,
}

const (
	blackIndex  = 1
	greenIndex  = 2
	redIndex    = 3
	yellowIndex = 4
	whiteIndex  = 5
)

func main() {
	http.HandleFunc("/count", counter)
	http.HandleFunc("/", handler)
	http.HandleFunc("/lissajous", lissaHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissaHandler(w http.ResponseWriter, r *http.Request) {
	cycles, err := strconv.ParseFloat(r.FormValue("c"), 64)

	if err != nil {
		cycles = 5
	}
	lissajous(w, cycles)
}

func lissajous(out io.Writer, cycles float64) {

	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 3
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			if t < 10 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), redIndex)
			} else if t < 20 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), whiteIndex)
			}

		}
		phase += 0.01
		if i+1 == nframes {
			anim.Delay = append(anim.Delay, 100)
		} else {
			anim.Delay = append(anim.Delay, delay)
		}

		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)

}
