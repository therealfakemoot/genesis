package render

import (
	"encoding/json"
	"io"
	"math"

	log "github.com/sirupsen/logrus"

	geo "github.com/therealfakemoot/genesis/geo"
)

func ServeJSON(w io.Writer, m geo.Map) {
	type mapData struct {
		Width  int       `json:"width"`
		Height int       `json:"height"`
		Seed   int       `json:"seed"`
		Steps  float64   `json:"steps"`
		Min    float64   `json:"min"`
		Max    float64   `json:"max"`
		Values []float64 `json:"values"`
	}
	var md mapData
	md.Width, md.Height = m.Width, m.Height
	md.Min, md.Max = m.Domain.Min, m.Domain.Max
	md.Seed = m.Seed
	md.Steps = (math.Abs(m.Domain.Min) + math.Abs(m.Domain.Min)) / 35

	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			v := m.Points[i][j]
			if v < 0.0 {
				v = 0.0
			}
			md.Values = append(md.Values, v)
		}
	}

	err := json.NewEncoder(w).Encode(md)
	if err != nil {
		log.WithError(err).Error("error sending map json")
	}
}
