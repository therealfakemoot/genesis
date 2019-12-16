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

func (m Map) Dims() (int, int) {
	return m.Width, m.Height
}

func (m Map) Z(c, r int) float64 {
	return m.Points[c][r]
}

func (m Map) X(c int) float64 {
	if c > m.Width {
		panic("out of bounds")
	}
	return float64(c)
}

func (m Map) Y(r int) float64 {
	if r > m.Height {
		panic("out of bounds")
	}
	return float64(r)
}

func (m Map) Min() float64 {
	return m.Domain.Min
}

func (m Map) Max() float64 {
	return m.Domain.Max
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
