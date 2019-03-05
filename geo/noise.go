package geo

import (
	noise "github.com/ojrac/opensimplex-go"

	Q "github.com/therealfakemoot/go-quantize"
)

type NoiseOpts struct {
}

func Noise(m Map) (points [][]float64) {
	return NoiseComplex(m)
}

func NoiseComplex(m Map) (points [][]float64) {
	// This is important. Adding the noise values together means the input domain grows.
	input := Q.Domain{Min: -3, Max: 3}
	n := noise.New(int64(m.Seed))

	x, y := m.Width, m.Height

	for i := 0; i < y; i++ {
		row := make([]float64, x)
		for j := 0; j < x; j++ {
			v := n.Eval2(float64(j)*0.001, float64(i)*0.001) +
				n.Eval2(float64(j)*0.05, float64(i)*0.05) +
				(0.25 * n.Eval2(float64(j)*0.1, float64(i)*0.1))

			row[j] = v
		}
		quantized := Q.QuantizeAll(row, input, m.Domain)

		points = append(points, quantized)
	}
	return points
}
