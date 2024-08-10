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

	flag.StringVar(&port, "port", "8888", "server port to bind to")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Throttle(5))

	r.Get("/dashboard", ServeDashboard)

	r.Group(func(r chi.Router) {
		r.Use(MapCtx)
		r.Get("/map/{target}", ServeMap)
	})

	log.Fatal(http.ListenAndServe(":"+port, r))
}
