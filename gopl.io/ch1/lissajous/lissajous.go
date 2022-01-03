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

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
}

const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
)

var parameters = map[string]float64{
	"cycles":  5,
	"res":     0.001,
	"size":    100,
	"nframes": 128,
	"delay":   8,
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	//lissajous(os.Stdout)
	if file, err := os.Create("out.gif"); err != nil {
		panic(err)
	} else {
		lissajous(file)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse: %v\n", err)
	}
	for k, v := range r.Form {
		value, _ := strconv.Atoi(v[0])
		parameters[k] = float64(value)
	}
	lissajous(w)
}

func lissajous(out io.Writer) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: int(parameters["nframes"])}
	phase := 0.0
	for i := 0; i < int(parameters["nframes"]); i++ {
		rect := image.Rect(0, 0, 2*int(parameters["size"])+1, 2*int(parameters["size"])+1)
		img := image.NewPaletted(rect, palette)
		colorIdx := uint8(rand.Intn(8))
		for t := 0.0; t < parameters["cycles"]*2*math.Pi; t += parameters["res"] {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(parameters["size"])+int(x*parameters["size"]+0.5), int(parameters["size"])+int(y*parameters["size"]+0.5), colorIdx)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, int(parameters["delay"]))
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
