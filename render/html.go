package render

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"

	geo "github.com/therealfakemoot/genesis/geo"
)

var RootTemplate = template.New("root")
var D3Template = template.Must(RootTemplate.ParseFiles("static/d3.tpl"))
var PlotlyTemplate = template.Must(RootTemplate.ParseFiles("static/plotly.tpl"))

func Plotly(w http.ResponseWriter, m geo.Map) {
	w.Header().Set("Content-Type", "text/html")
	err := RootTemplate.ExecuteTemplate(w, "plotly", m)
	if err != nil {
		log.WithError(err).Error("template execution")
	}
}

func D3(w http.ResponseWriter, m geo.Map) {
	w.Header().Set("Content-Type", "text/html")
	err := RootTemplate.ExecuteTemplate(w, "d3", m)
	if err != nil {
		log.WithError(err).Error("template execution")
	}
}
