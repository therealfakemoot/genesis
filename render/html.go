package render

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"

	geo "github.com/therealfakemoot/genesis/geo"
)

var TopoTemplate = template.Must(template.New("topo.tpl").ParseFiles("static/topo.tpl"))

func ServeHTML(w http.ResponseWriter, r *http.Request) {
	var m geo.Map
	w.Header().Set("Content-Type", "text/html")
	err := TopoTemplate.Execute(w, m)
	if err != nil {
		log.WithError(err).Error("template execution")
	}
}
