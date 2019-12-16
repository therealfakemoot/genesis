{{ define "colorPicker" }}
<form class="pure-form pure-form-aligned {{ . }}">
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
{{ end }}
