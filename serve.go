package main

import (
	"bytes"
	"encoding/json"
	"image/png"
	"net/http"
	"strconv"

	Q "github.com/therealfakemoot/go-quantize"

	log "github.com/sirupsen/logrus"
)

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

	d := Q.Domain{
		Max: float64(max),
		Min: float64(min),
	}

	m := GenerateMap(int(width), int(height), int(seed), d)
	m.Domain = d
	i := GeneratePNG(m)

	buffer := new(bytes.Buffer)

	w.Header().Set("Content-type", "image/png")

	w.Header().Set("Content-Disposition", `inline;filename="`+values.Encode()+`"`)
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

	log.WithFields(log.Fields{
		"seed":   seed,
		"width":  width,
		"height": height,
		"min":    min,
		"max":    max,
	}).Info("serving png")
}
