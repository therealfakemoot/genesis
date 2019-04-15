package main

import (
	"net/http"
	"os"
	// "flag"

	"github.com/go-chi/chi"
	chi_mid "github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"

	middleware "github.com/therealfakemoot/genesis/middleware"
	render "github.com/therealfakemoot/genesis/render"
)

func main() {
	r := chi.NewRouter()
	r.Use(chi_mid.RequestID)
	requestLogs, _ := os.OpenFile("genesis.requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	r.Use(middleware.RequestLogger(requestLogs, log.DebugLevel, false))
	r.Use(middleware.ClientLogger(os.Stdout, log.DebugLevel, false))
	r.Use(chi_mid.Recoverer)

	r.Group(func(r chi.Router) {
		r.Use(middleware.MapCtx)

		r.Post("/map/png/", render.ServePNG)

	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
