package main

import (
	"flag"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		port string
	)

	flag.StringVar(&port, "port", "8080", "server port to bind to")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Group(func(r chi.Router) {
		r.Use(MapCtx)
		r.Use(middleware.Throttle(5))

		r.Get("/map/{target}", ServeMap)

	})

	log.Fatal(http.ListenAndServe(":"+port, r))
}
