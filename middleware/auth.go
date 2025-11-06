package middleware

import (
	"net/http"
	"small_imgbed/internal/auth"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(401)
			return
		}
		if auth.Auth(authHeader) {
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(403)
	})
}
