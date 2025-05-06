package metrics

import (
	"net/http"

	"github.com/Furrity/web-resume/pkg/app"
	"github.com/go-chi/chi/v5"
)

func MetricsRouter(a *app.App) http.Handler {
	r := chi.NewRouter()

	r.Get("/", app.Adapter(a, MetricsHandler))

	return r
}
