package handler

import (
	"net/http"
	"small_imgbed/internal/storage"
	"strings"
)

func Get() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if strings.HasPrefix(p, "/") {
			w.Write(storage.Get(p[1:]))
			return
		}
		w.WriteHeader(400)
	})
	return router
}
