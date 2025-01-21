package middleware

import "net/http"

func AuthHandler(h http.Handler, log logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug("authorization completed")

		h.ServeHTTP(w, r)
	})
}
