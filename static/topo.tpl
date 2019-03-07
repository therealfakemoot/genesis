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

plot = function(terrain) {
	let thresholds = d3.range({{ $.Domain.Min }}, {{ $.Domain.Max }} + 1, terrain.steps)

	let contours = d3.contours()
	.size([terrain.width, terrain.height])
	.thresholds(thresholds)(terrain.values)

	svg.selectAll("path")
	.data(contours)
	.enter().append("path")
	.attr("d", d3.geoPath(d3.geoIdentity().scale(width / terrain.width)))
	.attr("fill", function(d) { return color(d.value); });

}

fetch("/json", {
body: JSON.stringify({
		      "width": {{ $.Width }},
		      "height": {{ $.Height }},
		      "seed": {{ $.Seed }},
		      "domain": {
		      "min": {{ $.Domain.Min }},
		      "max": {{ $.Domain.Max }},
		      }
		      })
}).then(response => {
	plot(response.json())
})

/*
d3.json("/json?width={{ $.Width }}&height={{ $.Height }}&seed={{ $.Seed }}&min={{ $.Domain.Min }}&max={{ $.Domain.Max }}&out=json", function(error, terrain) {
	if (error) {
		throw error;
	}
	plot(terrain)

});
*/

</script>
</body>
