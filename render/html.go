package render

import (
	"html/template"
	"log"
	"net/http"

	geo "github.com/therealfakemoot/genesis/geo"
)

var TopoTemplate = template.Must(template.New("terrain").ParseFiles("static/topo.tpl"))

func ServeHTML(w http.ResponseWriter, m geo.Map) {
	w.Header().Set("Content-Type", "text/html")
	log.Println(TopoTemplate)
	TopoTemplate.Execute(w, m)
}
