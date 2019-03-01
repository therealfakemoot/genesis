package geo

import (
	noise "github.com/ojrac/opensimplex-go"

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

	input := Q.Domain{Min: -1, Max: 1}
	n := noise.New(int64(seed))

	for i := 0; i < y; i++ {
		row := make([]float64, x)
		for j := 0; j < x; j++ {
			row[j] = n.Eval2(float64(j)*00.1, float64(i)*00.1)
		}
		quantized := Q.QuantizeAll(row, input, d)

		m.Points[i] = quantized
	}

	return m
}
