package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {

	http.HandleFunc("/map", ServePNG)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
