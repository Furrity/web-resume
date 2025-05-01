package handlers

import (
	"net/http"

	"github.com/Furrity/web-resume/internal"
	"github.com/Furrity/web-resume/pkg/app"
)

type MetricsResponse struct {
	Requests int `json:"requests"`
}

func MetricsHandler(a *app.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		internal.RespondWithJson(w, 200, MetricsResponse{Requests: a.GetRequestCount()})
	}
}
