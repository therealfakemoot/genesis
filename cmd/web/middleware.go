package main

import (
	"context"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	geo "github.com/therealfakemoot/genesis/geo"
	Q "github.com/therealfakemoot/go-quantize"
)

type CtxKey string

var (
	CtxMap    = CtxKey("map")
	CtxLogger = CtxKey("clientLogger")
)

func ClientLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

}

func MapCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var m geo.Map

		parseInt := func(s string) int {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return 0
			}
			return int(i)
		}

		seed := parseInt(r.FormValue("seed"))
		width := parseInt(r.FormValue("width"))
		height := parseInt(r.FormValue("height"))
		min := parseInt(r.FormValue("min"))
		max := parseInt(r.FormValue("max"))

		m.Seed = seed
		m.Width = width
		m.Height = height
		m.Domain = Q.Domain{
			Min: float64(min),
			Max: float64(max),
		}

		log.WithFields(log.Fields{
			"seed":   seed,
			"width":  width,
			"height": height,
			"min":    min,
			"max":    max,
		}).Info("raw url param values")

		valid := func(v bool) string {
			if v {
				return "invalid"
			}
			return "valid"
		}

		if (m.Width <= 0 || m.Height <= 0) || m.Domain.Min > m.Domain.Max {
			w.WriteHeader(400)
			log.WithFields(log.Fields{
				"width":  valid(m.Width <= 0),
				"height": valid(m.Height <= 0),
				"domain": valid(m.Domain.Min > m.Domain.Max),
			}).Error("invalid map")
			return

		}

		ctx := context.WithValue(r.Context(), CtxMap, m)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
