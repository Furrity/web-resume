package profile

import (
	"net/http"

	"github.com/Furrity/web-resume/pkg/app"
	"github.com/go-chi/chi/v5"
)

func ProfileRouter(a *app.App) http.Handler {
	r := chi.NewRouter()

	r.Get("/{lang)}", app.Adapter(a, GetProfileHandler))

	return r
}
