package geo

import (
	noise "github.com/ojrac/opensimplex-go"

	Q "github.com/therealfakemoot/go-quantize"
)

type NoiseOpts struct {
	Alpha []float64 `json:"alpha"`
}

var (
	NoiseDefaults = NoiseOpts{
		Alpha: []float64{0.001, 0.05, 0.1},
	}
)

func Noise(m Map) (points [][]float64) {
	return NoiseComplex(m, NoiseDefaults)
}

func NoiseComplex(m Map, no NoiseOpts) (points [][]float64) {
	// This is important. Adding the noise values together means the input domain grows.
	input := Q.Domain{Min: -3, Max: 3}
	n := noise.New(int64(m.Seed))

	x, y := m.Width, m.Height
	alphaFine, alphaMid, alphaCoarse := no.Alpha[0], no.Alpha[1], no.Alpha[2]

	for i := 0; i < y; i++ {
		row := make([]float64, x)
		for j := 0; j < x; j++ {
			v := n.Eval2(float64(j)*alphaFine, float64(i)*alphaFine) +
				n.Eval2(float64(j)*alphaMid, float64(i)*alphaMid) +
				(0.25 * n.Eval2(float64(j)*alphaCoarse, float64(i)*alphaCoarse))

			row[j] = v
		}
		quantized := Q.QuantizeAll(row, input, m.Domain)

		points = append(points, quantized)
	}
	return points
}
