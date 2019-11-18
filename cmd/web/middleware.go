package main

import (
	"context"
	"encoding/json"
	"net/http"

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
		decoder := json.NewDecoder(r.Body)
		var m geo.Map
		err := decoder.Decode(&m)
		if err != nil {
			// log.WithError(err).Error("error deserializing map")
			m.Width = 1000
			m.Height = 1000
			m.Seed = 123456
			m.Domain = Q.Domain{
				Min: -1000,
				Max: 1000,
			}
			ctx := context.WithValue(r.Context(), CtxMap, m)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

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
