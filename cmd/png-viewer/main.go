package main

import (
	"flag"
	"image/png"
	"log"
	"os"

	"github.com/therealfakemoot/genesis/geo"
	"github.com/therealfakemoot/genesis/render"
	Q "github.com/therealfakemoot/go-quantize"
)

func main() {
	var (
		x, y, seed int64
		min, max   int
		dump, load string
		m          geo.Map
	)

	flag.IntVar(&min, "min", -1000, "post-quantize minimum")
	flag.IntVar(&max, "max", 1000, "post-quantize maximum")
	flag.Int64Var(&x, "x", 1000, "map width")
	flag.Int64Var(&y, "y", 1000, "map height")
	flag.Int64Var(&seed, "seed", 42069, "simplex noise seed")
	flag.StringVar(&dump, "dump", "raw.map", "path to destination map file")
	flag.StringVar(&dump, "load", "raw.map", "path to map file to load")

	flag.Parse()

	if (dump != "") && (load != "") {
		log.Fatalf("dump and load are mutually exclusive")
	}

	if dump != "" {
		d := Q.Domain{Min: float64(min), Max: float64(max)}
		m = geo.New(x, y, seed, d)
		f, err := os.OpenFile(dump, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalf("couldn't open map file for dump: %s", err)
		}
		err = m.Pack(f)
		if err != nil {
			log.Fatalf("couldn't pack map: %s", err)
		}

	}

	if load != "" {
		f, err := os.Open(load)
		if err != nil {
			log.Fatalf("couldn't open map file for load: %s", err)
		}
		m, err = geo.Unpack(f)
		if err != nil {
			log.Fatalf("couldn't unpack map file: %s", err)
		}

	}

	i := render.GeneratePNG(m)

	err := png.Encode(os.Stdout, i)
	if err != nil {
		log.Fatalf("couldn't write PNG")
	}
}
