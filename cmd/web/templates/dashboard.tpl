<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8" />
		<title>{{ .Title }}</title>
	</head>
	<body>
		<h1>Noisemap Visualizer</h1>
		<form>
		<fieldset>
		<legend>Perlin Parameters</legend>
		<input type="range" name="octaves" min="1" max="9" list="octaves"/><label for="octaves">Octaves</label>
		</fieldset>
		</form>
		<form>
		<fieldset>
		<legend>Contour Map Parameters</legend>
		<input type="number" name="height_step" min="0"/><label for="height_step">Height Step</label>
		</fieldset>
		</form>
	</body>

<datalist id="octaves">
	<option value="1" />
	<option value="2" />
	<option value="3" />
	<option value="4" />
	<option value="5" />
	<option value="6" />
	<option value="7" />
	<option value="8" />
	<option value="9" />
</datalist>

</html>
