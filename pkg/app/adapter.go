package app

import "net/http"

type AppHandler func(http.ResponseWriter, *http.Request, *App)

func Adapter(app *App, handler AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, app)
	}
}
