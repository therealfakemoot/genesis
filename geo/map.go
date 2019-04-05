package geo

import (
	log "github.com/sirupsen/logrus"

	Q "github.com/therealfakemoot/go-quantize"
)

// Map represents a geographic region that is a 2d plane.
//
// A Map is the core construct for Genesis. All other values,
// can be derived from or must be saved to the context of a Map.
type Map struct {
	Seed   int         `json:"seed"`
	Domain Q.Domain    `json:"domain"`
	Width  int         `json:"width"`
	Height int         `json:"height"`
	Points [][]float64 `json:"points"`
	Log    *log.Logger
}

// Dims returns the X,Y lengths of the Map.
//
// For compliance with the gonum/plot/plotter.GridXYZ interface.
func (m Map) Dims() (int, int) {
	return m.Width, m.Height
}

// Z returns the height value at point (C, R).
func (m Map) Z(c, r int) float64 {
	return m.Points[c][r]
}

// X returns the X coordinate of column c.
func (m Map) X(c int) float64 {
	// The interface demands this. I think it's only useful
	// if your underlying data structure is sparse.
	if c > m.Width {
		panic("out of bounds")
	}
	return float64(c)
}

// Y returns the Y coordinate of row r.
func (m Map) Y(r int) float64 {
	// The interface demands this. I think it's only useful
	// if your underlying data structure is sparse.
	if r > m.Height {
		panic("out of bounds")
	}
	return float64(r)
}

// Min returns the smallest value in this instances' data set.
func (m Map) Min() float64 {
	// Implementing this to save on iterations over the set of all points.
	return m.Domain.Min
}

// Max returns the largest value in this instances' data set.
func (m Map) Max() float64 {
	// Implementing this to save on iterations over the set of all points.
	return m.Domain.Max
}

// New returns a new Map, pre-populated with terrain data.
func New(x, y, seed int, d Q.Domain) (m Map) {
	m.Seed = seed
	m.Points = make([][]float64, y)
	m.Width = x
	m.Height = y
	m.Domain = d

	m.Noise()

	return m
}
