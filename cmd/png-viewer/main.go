package main

import (

    "github.com/therealfakemoot/genesis/render"
    log "github.com/sirupsen/logrus"
    "github.com/therealfakemoot/genesis/geo"
    Q "github.com/therealfakemoot/go-quantize"
)

func main() {
    d := Q.Domain{Min: -1000, Max: 1000}
    m := geo.New(1000, 1000, 42069, d)
    i := render.GeneratePNG(m)
}
