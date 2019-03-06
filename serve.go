package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"

	geo "github.com/therealfakemoot/genesis/geo"
	render "github.com/therealfakemoot/genesis/render"
)

type CtxKey string

var (
	CtxMapKey = CtxKey("map")
)

func MapCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var m geo.Map
		err := decoder.Decode(&m)
		if err != nil {
			log.WithError(err).Error("error deserializing map")
			return
		}
		ctx := context.WithValue(r.Context(), CtxMapKey, m)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ServeMap(w http.ResponseWriter, r *http.Request) {
	out := chi.URLParam(r, "target")

	m := r.Context().Value(CtxMapKey).(geo.Map)

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

	switch out {
	case "png":
		render.ServePNG(w, m)
	case "json":
		w.Header().Set("Content-Type", "application/json")
		render.ServeJSON(w, m)
	case "html":
		render.ServeHTML(w, m)
	default:
		render.ServePNG(w, m)
	}

	mapCtx.Info("serving map")
}
