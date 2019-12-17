package geo

import (
	// log "github.com/sirupsen/logrus"

	Q "github.com/therealfakemoot/go-quantize"
)

// Map represents all data related to the geography, climate,
// and points of interest in a given world.
type Map struct {
	Seed   int         `json:"seed"`
	Domain Q.Domain    `json:"domain"`
	Width  int         `json:"width"`
	Height int         `json:"height"`
	Points [][]float64 `json:"points"`
}

// New returns a zeroed Map value.
//
// x, y: Describe the width and height of the new map.
// d   : The domain describes the largest and smallest values which can be generated in the context of this map.
// no  : The NoiseOpts value provides fine tuning constraints on the generated noise values; this gives multiple "knobs" for controlling how rough/smooth the noise space is.
func New(x, y, seed int, d Q.Domain, no NoiseOpts) (m Map) {
	m.Seed = seed
	m.Points = make([][]float64, y)
	m.Width = x
	m.Height = y
	m.Domain = d

	m.Points = NoiseComplex(m, no)

	return m
}
