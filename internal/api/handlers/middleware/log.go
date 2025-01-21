package middleware

import "net/http"

func LogHandler(h http.Handler, log loggerHTTP) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.DebugHTTP(r.URL.String(), r.Method)

		h.ServeHTTP(w, r)
	})
}
