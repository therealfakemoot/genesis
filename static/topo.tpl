<!DOCTYPE html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://unpkg.com/purecss@1.0.1/build/pure-min.css" integrity="sha384-oAOxQR6DkCoMliIh8yFnu25d7Eq/PHS21PClpwjOTeU2jRSq11vu66rf90/cZr47" crossorigin="anonymous">
<script src="https://cdn.plot.ly/plotly-latest.min.js"></script>
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
			<label for="floor">
				<input type="color" name="floor" />
				Floor
			</label>
			<label for="ceiling">
				<input type="color" name="ceiling" />
				Ceiling
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
	<svg class = "pure-u-1" width="{{ $.Width }}" height="{{ $.Height }}" stroke="#fff" stroke-width="0.5"></svg>
</div>

<script src="https://d3js.org/d3.v4.min.js"></script>
<script src="https://d3js.org/d3-hsv.v0.1.min.js"></script>
<script src="https://d3js.org/d3-contour.v1.min.js"></script>
<script>

var svg = d3.select("svg"),
width = +svg.attr("width"),
height = +svg.attr("height");

var i0 = d3.interpolateHsvLong(d3.hsv(120, 1, 0.65), d3.hsv(60, 1, 0.90)),
i1 = d3.interpolateHsvLong(d3.hsv(60,1, 0.90), d3.hsv(0, 0, 0.95)),
interpolateTerrain = function(t) { return t < 0.5 ? i0(t * 2) : i1((t - 0.5) * 2); },
color = d3.scaleSequential(interpolateTerrain).domain([{{ $.Domain.Min }}, {{ $.Domain.Max }}]);

d3.json("/map/json?width={{ $.Width }}&height={{ $.Height }}&seed={{ $.Seed }}&min={{ $.Domain.Min }}&max={{ $.Domain.Max }}", function(error, terrain) {
	if (error) {
		throw error;
	}

	let thresholds = d3.range({{ $.Domain.Min }}, {{ $.Domain.Max }} + 1, terrain.steps)

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
