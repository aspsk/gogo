package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"bufio"
	"math"
	"math/rand"
	"strconv"
	"net/http"
)

type respondParams = struct {
	cycles int
	res float64
	size int
	nframes int
	delay int
}

func newRespondParams() *respondParams {
	ret := respondParams{
		cycles: 5,
		res: 0.001,
		size: 100,
		nframes: 64,
		delay: 8,
	}
	return &ret
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func randColor() color.Color {

	x := rand.Uint32()
	r := byte(x & 0xff)
	g := byte((x & 0xff00) >> 8)
	b := byte((x & 0xff0000) >> 16)
	return color.RGBA{r, g, b, 0xff}

}

func make_palette() []color.Color {
	return []color.Color{color.Black, color.RGBA{0x00, 0x88, 0x00, 0xff}}
}

func lissajous(out io.Writer, make_palette func() []color.Color, params *respondParams) {

	size := params.size
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: params.nframes}
	phase := 0.0
	for i := 0; i < params.nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, make_palette())
		for t := 0.0; t < float64(params.cycles)*2*math.Pi; t += params.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size + int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, params.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}

func respond(w http.ResponseWriter, params *respondParams) {
	var b bytes.Buffer
	buf := bufio.NewWriter(&b)

	lissajous(buf, make_palette, params)

	r := bufio.NewReader(&b)
	if _, err := io.Copy(w, r); err != nil {
		fmt.Println(err)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := newRespondParams()

	if x := r.FormValue("cycles"); x != "" {
		cycles, err := strconv.Atoi(x)
		if err == nil {
			params.cycles = cycles
		}
	}

	if x := r.FormValue("res"); x != "" {
		res, err := strconv.ParseFloat(x, 64)
		if err == nil {
			params.res = res
		}
	}

	if x := r.FormValue("size"); x != "" {
		size, err := strconv.Atoi(x)
		if err == nil {
			params.size = size
		}
	}

	if x := r.FormValue("nframes"); x != "" {
		nframes, err := strconv.Atoi(x)
		if err == nil {
			params.nframes = nframes
		}
	}

	if x := r.FormValue("delay"); x != "" {
		delay, err := strconv.Atoi(x)
		if err == nil {
			params.delay = delay
		}
	}

	respond(w, params)
}

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
