{{ define "mapControls" }}
<div class="pure-g {{ . }}">
	{{ template "colorPicker" "" }}
	{{ template "mapParams" "" }}
</div>
{{ end }}
