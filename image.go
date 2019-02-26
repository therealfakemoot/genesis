package main

import (
	"image"
	"image/color"

	Q "github.com/therealfakemoot/go-quantize"

	log "github.com/sirupsen/logrus"
)

func matchColor(point float64, d Q.Domain) (c color.Color) {
	colorSpace := Q.Domain{
		Min: 0,
		Max: 255,
	}
	// normalized := uint8(colorSpace.QuantizePoint(point))
	normalized := uint8(Q.Quantize(point, d, colorSpace))
	log.WithFields(log.Fields{
		"point": point,
		"color": normalized,
	}).Debug("matching color")

	return color.NRGBA{normalized, normalized, normalized, 255}
}

func GeneratePNG(m Map) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, m.Width, m.Height))

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			point := m.Points[x][y]
			c := matchColor(point, m.Domain)
			img.Set(x, y, c)
		}
	}

	return img
}
