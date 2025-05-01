package middleware

import "net/http"

func IncrementStatsMiddleware(incrementStatsFunc func()) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			incrementStatsFunc()
			next.ServeHTTP(w, r)
		})
	}
}
