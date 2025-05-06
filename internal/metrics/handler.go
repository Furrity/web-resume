package metrics

import (
	"net/http"

	"github.com/Furrity/web-resume/pkg/app"
	"github.com/Furrity/web-resume/pkg/utils"
)

type MetricsResponse struct {
	Requests int `json:"requests"`
}

func MetricsHandler(w http.ResponseWriter, r *http.Request, a *app.App) {
	utils.RespondWithJson(w, 200, MetricsResponse{Requests: a.GetRequestCount()})
}
