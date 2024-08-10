package main

import (
	"embed"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"html/template"

	geo "github.com/therealfakemoot/genesis/geo"
	render "github.com/therealfakemoot/genesis/render"
)

func init() {
	rawDashboardTemplate, err := templates.ReadFile("templates/dashboard.tpl")
	if err != nil {
		log.Errorf("error fetching dashboard template from embedded fs: %s", err)
	}

	dashboardTemplate = template.Must(template.New("dashboard").Parse(string(rawDashboardTemplate)))
}

var (
	dashboardTemplate *template.Template
)

type DashboardTemplateData struct {
	Title string
}

//go:embed templates/*
var templates embed.FS

func ServeDashboard(w http.ResponseWriter, r *http.Request) {
	dtd := DashboardTemplateData{Title: "Perlin noisemap visualizer"}
	err := dashboardTemplate.Execute(w, dtd)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func ServeMap(w http.ResponseWriter, r *http.Request) {
	out := chi.URLParam(r, "target")

	m := r.Context().Value(CtxMap).(geo.Map)

	width, height := m.Width, m.Height
	seed := m.Seed
	d := m.Domain

	var noiseOpts geo.NoiseOpts
	_, err := toml.DecodeFile("noiseOpts.toml", &noiseOpts)
	if err != nil {
		log.WithError(err).Error("unable to open noiseOpts config file")
	}

	m = geo.New(int(width), int(height), int(seed), d, noiseOpts)

	mapCtx := log.WithFields(log.Fields{
		"target": out,
		"seed":   seed,
		"width":  width,
		"height": height,
		"min":    d.Min,
		"max":    d.Max,
	})

	switch out {
	case "png":
		render.ServePNG(w, m)
	case "json":
		w.Header().Set("Content-Type", "application/json")
		render.ServeJSON(w, m)
	case "d3":
		render.D3(w, m)
	case "plotly":
		render.Plotly(w, m)
	default:
		render.ServePNG(w, m)
	}

	mapCtx.Info("serving map")
}
