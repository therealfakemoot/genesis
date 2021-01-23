package main

import (
    "os"
    "image/png"
    "flag"

    "github.com/therealfakemoot/genesis/render"
    "github.com/therealfakemoot/genesis/geo"
    Q "github.com/therealfakemoot/go-quantize"
)

func main() {
    var (
        min, max, x, y, seed int
    )

    flag.IntVar(&min, "min",  -1000, "post-quantize minimum")
    flag.IntVar(&max, "max", 1000, "post-quantize maximum")
    flag.IntVar(&x, "x", 1000, "map width")
    flag.IntVar(&y, "y", 1000, "map height")
    flag.IntVar(&seed, "seed", 42069, "simplex noise seed")

    d := Q.Domain{Min: float64(min), Max: float64(max)}
    m := geo.New(x, y, seed, d)
    // i := render.GeneratePNG(m)



    png.Encode(os.Stdout, i)
}
