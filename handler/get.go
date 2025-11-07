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
		if !strings.HasPrefix(p, "/") {
			w.WriteHeader(400)
			return
		}
		if !storage.FileExit(p[1:]) {
			w.WriteHeader(404)
			return
		}
		w.Write(storage.Get(p[1:]))
	})
	return router
}
