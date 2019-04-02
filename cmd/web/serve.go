package main

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func ServeReact(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("public/index.html")
	if err != nil {
		w.WriteHeader(404)
		log.WithError(err).Error("couldn't load index file")
	}
	io.Copy(w, f)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
