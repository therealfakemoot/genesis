{{ define "mapParams" }}
<div class="{{ . }}">
<form class="pure-form pure-form-aligned">
	<fieldset class="pure-group">
		<legend>Map Parameters</legend>
		<div class="pure-control-group">
			<label for="width">Map Width</label>
			<input class="pure-input-rounded" type="number" name="width" value="1000"/>
		</div>
		<div class="pure-control-group">
			<label for="height">Map Height</label>
			<input class="pure-input-rounded" type="number" name="height" value="1000"/>
		</div>
	</fieldset>
</form>
</div>
{{ end }}
