package main

import (
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"

	geo "github.com/therealfakemoot/genesis/geo"
	middleware "github.com/therealfakemoot/genesis/middleware"
)

func ServeMap(w http.ResponseWriter, r *http.Request) {
	out := chi.URLParam(r, "target")

	m := r.Context().Value(middleware.CtxMap).(geo.Map)

	width, height := m.Width, m.Height
	seed := m.Seed
	d := m.Domain

	m = geo.New(int(width), int(height), int(seed), d)

	mapCtx := log.WithFields(log.Fields{
		"target": out,
		"seed":   seed,
		"width":  width,
		"height": height,
		"min":    d.Min,
		"max":    d.Max,
	})

	mapCtx.Info("serving map")
}
