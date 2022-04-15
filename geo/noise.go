package geo

import (
	noise "github.com/ojrac/opensimplex-go"
	log "github.com/sirupsen/logrus"

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

func (m *Map) Noise() {
	m.NoiseComplex(NoiseDefaults)
}

func (m *Map) NoiseComplex(no NoiseOpts) {
	// This is important. Adding the noise values together means the input domain grows.
	input := Q.Domain{Min: -3, Max: 3}
	m.Log.Debug("entering NoiseComplex")
	n := noise.New(int64(m.Seed))

	var points [][]float64
	x, y := m.Width, m.Height

	m.Log.WithFields(log.Fields{
		"seed":   m.Seed,
		"width":  x,
		"height": y,
	}).Debug("preparing map")

	alphaFine, alphaMid, alphaCoarse := no.Alpha[0], no.Alpha[1], no.Alpha[2]
	m.Log.WithFields(log.Fields{
		"fine":   alphaFine,
		"mid":    alphaMid,
		"coarse": alphaCoarse,
	}).Debug("alpha coefficients loaded")

	for i := int64(0); i < y; i++ {
		row := make([]float64, x)
		for j := int64(0); j < x; j++ {
			v := n.Eval2(float64(j)*alphaFine, float64(i)*alphaFine) +
				n.Eval2(float64(j)*alphaMid, float64(i)*alphaMid) +
				(0.25 * n.Eval2(float64(j)*alphaCoarse, float64(i)*alphaCoarse))

			row[j] = v
		}
		quantized := Q.QuantizeAll(row, input, m.Domain)

		points = append(points, quantized)
	}
	m.Points = points
}
