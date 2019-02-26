package main

import (
	"bytes"
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"strconv"

	noise "github.com/ojrac/opensimplex-go"
	Q "github.com/therealfakemoot/go-quantize"

	log "github.com/sirupsen/logrus"
)

type Map struct {
	Domain Q.Domain
	Width  int
	Height int
	Points [][]float64
}
type parseError struct {
	Param string
	Error string
}

func ServePNG(w http.ResponseWriter, r *http.Request) {
	var errors []parseError

	values := r.URL.Query()
	width, err := strconv.ParseInt(values.Get("width"), 10, 0)
	if err != nil {
		errors = append(errors, parseError{"width", err.Error()})
	}
	height, err := strconv.ParseInt(values.Get("height"), 10, 0)
	if err != nil {
		errors = append(errors, parseError{"height", err.Error()})
	}
	seed, err := strconv.ParseInt(values.Get("seed"), 10, 0)
	if err != nil {
		errors = append(errors, parseError{"seed", err.Error()})
	}
	min, err := strconv.ParseInt(values.Get("min"), 10, 0)
	if err != nil {
		errors = append(errors, parseError{"min", err.Error()})
	}
	max, err := strconv.ParseInt(values.Get("max"), 10, 0)
	if err != nil {
		errors = append(errors, parseError{"max", err.Error()})
	}

	if len(errors) > 0 {
		log.WithFields(log.Fields{
			"errors": errors,
		}).Error("parsing parameters failed")

		e := struct {
			Errors []parseError
		}{errors}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(e)
		return
	}

	log.WithFields(log.Fields{
		"seed":   seed,
		"width":  width,
		"height": height,
		"min":    min,
		"max":    max,
	}).Info("serving png")

	d := Q.Domain{
		Max: float64(max),
		Min: float64(min),
	}

	m := GenerateMap(int(width), int(height), int(seed), d)
	m.Domain = d
	i := GeneratePNG(m)

	buffer := new(bytes.Buffer)

	w.Header().Set("Content-type", "image/png")
	err = png.Encode(buffer, i)
	if err != nil {
		log.WithError(err).Error("image encoding failure")

		e := struct {
			Error string
		}{Error: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(e)
		return
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	_, err = w.Write(buffer.Bytes())
	if err != nil {
		log.WithError(err).Error("response write failure")
	}
}

func matchColor(point float64, d Q.Domain) (c color.Color) {
	colorSpace := Q.Domain{
		Min: 0,
		Max: 255,
	}
	// normalized := uint8(colorSpace.QuantizePoint(point))
	normalized := uint8(Q.Quantize(point, d, colorSpace))
	log.WithFields(log.Fields{
		"point": point,
		"color": normalized,
	}).Debug("matching color")

	return color.NRGBA{normalized, normalized, normalized, 255}
}

func GeneratePNG(m Map) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, m.Width, m.Height))

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			point := m.Points[x][y]
			c := matchColor(point, m.Domain)
			img.Set(x, y, c)
		}
	}

	return img
}

func GenerateMap(x, y, seed int, d Q.Domain) (m Map) {
	m.Points = make([][]float64, y)
	m.Width = x
	m.Height = y

	input := Q.Domain{Min: -1, Max: 1}
	n := noise.New(int64(seed))

	for i := 0; i < y; i++ {
		row := make([]float64, x)
		for j := 0; j < x; j++ {
			row[j] = n.Eval2(float64(j), float64(i))
		}
		quantized := Q.QuantizeAll(row, input, d)

		m.Points[i] = quantized
	}

	return m
}

func main() {

	http.HandleFunc("/map", ServePNG)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
