package geo

import (
	// log "github.com/sirupsen/logrus"

	Q "github.com/therealfakemoot/go-quantize"
)

type Map struct {
	Seed   int         `json:"seed"`
	Domain Q.Domain    `json:"domain"`
	Width  int         `json:"width"`
	Height int         `json:"height"`
	Points [][]float64 `json:"points"`
}

func New(x, y, seed int, d Q.Domain, no NoiseOpts) (m Map) {
	m.Seed = seed
	m.Points = make([][]float64, y)
	m.Width = x
	m.Height = y
	m.Domain = d

	m.Points = NoiseComplex(m, no)

	return m
}
