{{ define "d3" }}

{{ template "header" }}

<body>


<div><h3>Topographical Terrain Map</h3></div>
<div class="pure-g">
	<div class="pure-u-1-2">
		<svg width="{{ $.Width }}" height="{{ $.Height }}" stroke="#fff" stroke-width="0.5"></svg>
	</div>
	<div class="pure-u-1-2">
		{{ template "mapControls" "pure-u-1-2"}}
	</div>
</div>

<script src="https://d3js.org/d3.v4.min.js"></script>
<script src="https://d3js.org/d3-hsv.v0.1.min.js"></script>
<script src="https://d3js.org/d3-contour.v1.min.js"></script>
<script>

var svg = d3.select("svg"),
width = +svg.attr("width"),
height = +svg.attr("height");

d3.json("/map/json?width={{ $.Width }}&height={{ $.Height }}&seed={{ $.Seed }}&min={{ $.Domain.Min }}&max={{ $.Domain.Max }}", function(error, terrain) {
	if (error) {
		throw error;
	}

	let thresholds = d3.range({{ $.Domain.Min }}, {{ $.Domain.Max }} + 1, terrain.steps)

	color = d3.scaleLinear()
	.domain([{{ $.Domain.Min }}, {{ $.Domain.Max }}])
	.range([0,1])
	.interpolate(d => d3.interpolatePlasma)

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

{{ template "footer" }}

</body>
{{ end }}
