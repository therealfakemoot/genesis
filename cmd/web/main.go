package main

import (
	"net/http"
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
	r.Use(chi_mid.Logger)
	r.Use(chi_mid.Recoverer)

	r.Get("/map/html", render.ServeHTML)

	r.Group(func(r chi.Router) {
		r.Use(middleware.MapCtx)
		// r.Use(chi_mid.Throttle(5))

		r.Post("/map/png", render.ServePNG)
		r.Post("/map/json", render.ServeJSON)

	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
