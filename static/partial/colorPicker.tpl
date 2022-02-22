{{ define "colorPicker" }}
<div class="{{ . }}">
<form class="pure-form pure-form-aligned">
	<fieldset class="pure-group">
		<legend>Map Colors</legend>
		<div class="pure-control-group">
			<label for="colorScheme">Color Scheme</label>
			<select name="colorScheme" size="4">
				<option value="Plasma" selected>Plasma</option>
				<option value="Turbo">Turbo</option>
				<option value="Viridis">Viridis</option>
				<option value="Inferno">Inferno</option>
				<option value="Magma">Magma</option>
				<option value="Warm">Warm</option>
				<option value="Cold">Cold</option>
			</select>
		</div>
	</fieldset>
</form>
</div>
{{ end }}
