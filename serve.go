package main

import (
	"encoding/json"
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

	switch out {
	case "png":
		render.ServePNG(w, m)
	case "json":
		w.Header().Set("Content-Type", "application/json")
		render.ServeJSON(w, m)
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
