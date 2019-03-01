package main

import (
	"net/http"
	"text/template"
)

var rawTemplate = `
<!DOCTYPE html>
<style>
body {
	background: black;
}
</style>
<body>
<svg width="{{ $.Width }}" height="{{ $.Height }}" stroke="#fff" stroke-width="0.5"></svg>
<script src="https://d3js.org/d3.v4.min.js"></script>
<script src="https://d3js.org/d3-hsv.v0.1.min.js"></script>
<script src="https://d3js.org/d3-contour.v1.min.js"></script>
<script>

var svg = d3.select("svg"),
width = +svg.attr("width"),
height = +svg.attr("height");

var i0 = d3.interpolateHsvLong(d3.hsv(120, 1, 0.65), d3.hsv(60, 1, 0.90)),
i1 = d3.interpolateHsvLong(d3.hsv(60, 1, 0.90), d3.hsv(0, 0, 0.95)),
interpolateTerrain = function(t) { return t < 0.5 ? i0(t * 2) : i1((t - 0.5) * 2); },
color = d3.scaleSequential(interpolateTerrain).domain([{{ $.Domain.Min }}, {{ $.Domain.Max }}]);

d3.json("/map?width={{ $.Width }}&height={{ $.Height }}&seed={{ $.Seed }}&min={{ $.Domain.Min }}&max={{ $.Domain.Max }}&out=json", function(error, terrain) {
	if (error) {
		throw error;
	}

	let thresholds = d3.range({{ $.Domain.Min }}, {{ $.Domain.Max }} + 1, 50)

	let contours = d3.contours()
	.size([terrain.width, terrain.height])
	.thresholds(thresholds)(terrain.values)

	svg.selectAll("path")
	.data(contours)
	.enter().append("path")
	.attr("d", d3.geoPath(d3.geoIdentity().scale(width / terrain.width)))
	.attr("fill", function(d) { return color(d.value); });
});

</script>
</body>
`

var TopoMap = template.Must(template.New("terrain").Parse(rawTemplate))

func ServeHTML(w http.ResponseWriter, m Map) {
	w.Header().Set("Content-Type", "text/html")
	TopoMap.Execute(w, m)
}
