{{ define "plotly" }}
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
	<h3>Topographical Terrain Map</h3>
</div>

</body>
{{ end }}
