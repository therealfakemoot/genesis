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

func New(x, y, seed int, d Q.Domain) (m Map) {
	m.Seed = seed
	m.Points = make([][]float64, y)
	m.Width = x
	m.Height = y

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
