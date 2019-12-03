{{ define "d3" }}
<!DOCTYPE html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://unpkg.com/purecss@1.0.1/build/pure-min.css" integrity="sha384-oAOxQR6DkCoMliIh8yFnu25d7Eq/PHS21PClpwjOTeU2jRSq11vu66rf90/cZr47" crossorigin="anonymous">
</head>

<style>
body {
	background: #4a306d;
	color: #000000;
}

.pure-form legend {
	color: #000000;
	border-bottom: .25em solid #B796AC;
}

</style>

<body>

<div class="pure-g">
	<form class="pure-form pure-u-1-3 pure-form-aligned">
		<fieldset class="pure-group">
			<legend>Map Colors</legend>
			<label for="colorScheme">
				<select name="colorScheme" size="4">
					<option value="Plasma" selected>Plasma</option>
					<option value="Turbo">Turbo</option>
					<option value="Viridis">Viridis</option>
					<option value="Inferno">Inferno</option>
					<option value="Magma">Magma</option>
					<option value="Warm">Warm</option>
					<option value="Cold">Cold</option>
				</select>
				Color Scheme
			</label>
		</fieldset>
	</form>

	<form class="pure-form pure-u-1-3 pure-form-aligned">
		<fieldset class="pure-group">
			<legend>Map Parameters</legend>
			<label for="width">
				<input class="pure-input-rounded" type="number" name="width" value="1000"/>
				Map Width
			</label>
			<label for="height">
				<input class="pure-input-rounded" type="number" name="height" value="1000"/>
				Map Height
			</label>
		</fieldset>
	</form>
</div>

<div class="pure-g">
	<h3>Topographical Terrain Map</h3>
	<svg class="pure-u-1" width="{{ $.Width }}" height="{{ $.Height }}" stroke="#fff" stroke-width="0.5"></svg>
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
</body>
{{ end }}
