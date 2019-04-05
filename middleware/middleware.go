package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	middleware "github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"

	geo "github.com/therealfakemoot/genesis/geo"
)

// CtxKey will prevent key collisions when passing values
// through middleware etc.
type CtxKey string

const (
	CtxMap           = CtxKey("map")
	CtxRequestLogger = CtxKey("clientLogger")
	CtxLogger        = CtxKey("logger")
)

func ClientLogger(out io.Writer, level log.Level, json bool) func(http.Handler) http.Handler {
	var formatter log.Formatter

	if json {
		formatter = new(log.JSONFormatter)
	} else {
		formatter = &log.TextFormatter{
			FullTimestamp: true,
		}
	}

	var logger = &log.Logger{
		Out:       out,
		Formatter: formatter,
		Hooks:     make(log.LevelHooks),
		Level:     level,
	}

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), CtxLogger, logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// ClientLogger configures a logrus context specifically designed to
// emit JSON errors.
func RequestLogger(out io.Writer, level log.Level, json bool) func(http.Handler) http.Handler {
	var formatter log.Formatter

	if json {
		formatter = new(log.JSONFormatter)
	} else {
		formatter = &log.TextFormatter{
			FullTimestamp: true,
		}
	}

	var logger = &log.Logger{
		Out:       out,
		Formatter: formatter,
		Hooks:     make(log.LevelHooks),
		Level:     level,
	}

	logger.Info("logger initialized")

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			base := logger.WithFields(log.Fields{
				"request-id": middleware.GetReqID(r.Context()),
				"url":        r.URL.Path,
				"method":     r.Method,
			})

			ctx := context.WithValue(r.Context(), CtxRequestLogger, base)
			start := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			defer func() {
				base.WithFields(log.Fields{
					"status":   ww.Status(),
					"size":     ww.BytesWritten(),
					"duration": time.Since(start),
				}).Debug(fmt.Sprintf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr))
			}()
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// MapCtx parses a geo.Map from the request body and injects it into
// the context.
func MapCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := r.Context().Value(CtxLogger).(*log.Logger)
		decoder := json.NewDecoder(r.Body)
		var m geo.Map
		m.Log = logger.WithField("request-id", middleware.GetReqID(r.Context()))
		err := decoder.Decode(&m)
		logger.WithFields(log.Fields{
			"map": fmt.Sprintf("%+v", m),
		}).Debug("decoded map")
		if err != nil {
			w.WriteHeader(400)
			logger.WithError(err).Error("error deserializing map")
			return
		}

		// TODO: Validate the map parameters

		ctx := context.WithValue(r.Context(), CtxMap, m)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
