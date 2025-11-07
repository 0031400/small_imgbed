package handler

import (
	"net/http"
)

func Login() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})
	return router
}
