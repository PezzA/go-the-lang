package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

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
	writer, _ := os.Create("output.gif")

	lissajous(writer)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
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
