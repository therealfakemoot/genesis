package geo

import (
	noise "github.com/ojrac/opensimplex-go"
	// log "github.com/sirupsen/logrus"

	Q "github.com/therealfakemoot/go-quantize"
)

type Map struct {
	Seed   int
	Domain Q.Domain
	Width  int
	Height int
	Points [][]float64
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

func New(x, y, seed int, d Q.Domain) (m Map) {
	m.Seed = seed
	m.Points = make([][]float64, y)
	m.Width = x
	m.Height = y
	m.Domain = d

	// This is important. Adding the noise values together means the input domain grows.
	input := Q.Domain{Min: -3, Max: 3}
	n := noise.New(int64(seed))

	for i := 0; i < y; i++ {
		row := make([]float64, x)
		for j := 0; j < x; j++ {
			v := n.Eval2(float64(j)*0.001, float64(i)*0.001) +
				n.Eval2(float64(j)*0.05, float64(i)*0.05) +
				(0.25 * n.Eval2(float64(j)*0.1, float64(i)*0.1))

			row[j] = v
		}
		quantized := Q.QuantizeAll(row, input, d)

		m.Points[i] = quantized
	}

	return m
}
