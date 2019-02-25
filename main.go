package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"

	"image"
	"image/color"
	"image/png"

	noise "github.com/ojrac/opensimplex-go"
	Q "github.com/therealfakemoot/go-quantize"
)

type Map struct {
	Width  int
	Height int
	Points [][]float64
}

func ServePNG(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	width, _ := strconv.ParseInt(values.Get("width"), 10, 0)
	height, _ := strconv.ParseInt(values.Get("height"), 10, 0)
	seed, _ := strconv.ParseInt(values.Get("seed"), 10, 0)
	min, _ := strconv.ParseInt(values.Get("min"), 10, 0)
	max, _ := strconv.ParseInt(values.Get("max"), 10, 0)

	log.Printf("seed: %d", seed)
	log.Printf("width: %d", width)
	log.Printf("height: %d", height)
	log.Printf("min: %d", min)
	log.Printf("max: %d", max)

	d := Q.Domain{
		Max:  float64(max),
		Min:  float64(min),
		Step: 1,
	}

	m := GenerateMap(500, 500, 18006665432, d)
	i := GeneratePNG(m)

	buffer := new(bytes.Buffer)

	w.Header().Set("Content-type", "image/png")
	err := png.Encode(buffer, i)
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if err != nil {
		log.Fatal(err)
	}
	w.Write(buffer.Bytes())
}

func GeneratePNG(m Map) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, m.Width, m.Height))

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			img.Set(x, y, color.NRGBA{0, 0, 0, 255})
		}
	}

	return img
}

func GenerateMap(x, y, seed int, d Q.Domain) (m Map) {
	m.Points = make([][]float64, y)
	m.Width = x
	m.Height = y

	n := noise.New(int64(seed))

	for yGen := 0; yGen < y; yGen++ {
		row := make([]float64, x)
		for xGen := 0; xGen < x; xGen++ {
			row[xGen] = n.Eval2(float64(xGen*2), float64(yGen*2))
		}
		quantized := d.Quantize(row)

		m.Points[yGen] = quantized
	}

	return m
}

func main() {

	http.HandleFunc("/map", ServePNG)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
