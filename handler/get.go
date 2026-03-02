package handler

import (
	"net/http"
	"regexp"
	"small_imgbed/internal/storage"
)

func Get() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if !isValid(p) {
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

var pattern = regexp.MustCompile(
	`^/\d{4}/\d{2}/\d{2}/[a-zA-Z0-9-]+\.[a-zA-Z0-9]+$`,
)

func isValid(s string) bool {
	return pattern.MatchString(s)
}
