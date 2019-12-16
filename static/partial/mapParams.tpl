{{ define "mapParams" }}
<form class="pure-form pure-form-aligned {{ . }}">
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
{{ end }}
