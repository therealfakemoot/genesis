package render

import (
	"html/template"
	"log"
	"net/http"

	log "github.com/sirupsen/logrus"

	geo "github.com/therealfakemoot/genesis/geo"
)

var TopoTemplate = template.Must(template.New("topo.tpl").ParseFiles("static/topo.tpl"))

func ServeHTML(w http.ResponseWriter, m geo.Map) {
	w.Header().Set("Content-Type", "text/html")
	err := TopoTemplate.Execute(w, m)
	if err != nil {
		log.WithError(err).Error("template execution")
	}
}
