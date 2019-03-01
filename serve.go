package main

import (
	"bytes"
	"encoding/json"
	"image/png"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	geo "github.com/therealfakemoot/genesis/geo"
	render "github.com/therealfakemoot/genesis/render"
	Q "github.com/therealfakemoot/go-quantize"
)

type parseError struct {
	Param string
	Error string
}

func ServeJSON(w http.ResponseWriter, m geo.Map) {
	type mapData struct {
		Width  int       `json:"width"`
		Height int       `json:"height"`
		Seed   int       `json:"seed"`
		Min    float64   `json:"min"`
		Max    float64   `json:"max"`
		Values []float64 `json:"values"`
	}
	var md mapData
	md.Width, md.Height = m.Width, m.Height
	md.Min, md.Max = m.Domain.Min, m.Domain.Max
	md.Seed = m.Seed

	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			v := m.Points[i][j]
			if v < 0.0 {
				v = 0.0
			}
			md.Values = append(md.Values, v)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(md)
	if err != nil {
		log.WithError(err).Error("error sending map json")
	}
}

func ServePNG(w http.ResponseWriter, m geo.Map) {
	var err error
	buffer := new(bytes.Buffer)

	i := render.GeneratePNG(m)

	w.Header().Set("Content-type", "image/png")

	w.Header().Set("Content-Disposition", `inline;filename="butts"`)
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

func ServeMap(w http.ResponseWriter, r *http.Request) {
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
	out := values.Get("out")
	if out == "" {
		out = "png"
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

	d := Q.Domain{
		Max: float64(max),
		Min: float64(min),
	}

	m := geo.New(int(width), int(height), int(seed), d)
	m.Domain = d

	switch out {
	case "png":
		ServePNG(w, m)
	case "json":
		ServeJSON(w, m)
	case "html":
		render.ServeHTML(w, m)
	}

	log.WithFields(log.Fields{
		"target": out,
		"seed":   seed,
		"width":  width,
		"height": height,
		"min":    min,
		"max":    max,
	}).Info("serving map")
}
