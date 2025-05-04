package handlers

import (
	"net/http"

	"github.com/Furrity/web-resume/pkg/app"
)

type AppHandler func(http.ResponseWriter, *http.Request, *app.App)

func WithApp(app *app.App, handler AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, app)
	}
}
