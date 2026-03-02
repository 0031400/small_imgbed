package handler

import (
	"net/http"
	"path/filepath"
	"small_imgbed/config"
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
		absPath, err := filepath.Abs(storage.GetPath(p[1:]))
		if err != nil {
			w.WriteHeader(500)
			return
		}
		absDataPath, err := filepath.Abs(config.C.Data.Path)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		if !strings.HasPrefix(absPath, absDataPath) {
			w.WriteHeader(400)
			return
		}
		exit, err := storage.FileExit(absPath)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		if !exit {
			w.WriteHeader(404)
			return
		}
		b, err := storage.Get(absPath)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		ext := filepath.Ext(absPath)
		if len(ext) > 0 {
			ext = ext[1:]
		}
		if mime, ok := config.C.Mime[ext]; ok {
			w.Header().Set("Content-Type", mime)
		}
		w.Write(b)
	})
	return router
}
