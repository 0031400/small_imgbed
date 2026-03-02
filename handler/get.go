package handler

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"small_imgbed/config"
	"small_imgbed/internal/storage"
	"strings"
)

func Get() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if p == "/" && config.C.Server.RootHtml {
			p = "/index.html"
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
			if config.C.Server.CopySite != "" {
				resp, err := http.Get(config.C.Server.CopySite + r.URL.Path)
				if err != nil {
					log.Println(err)
					w.WriteHeader(500)
					return
				}
				defer resp.Body.Close()
				if resp.StatusCode == 404 {
					w.WriteHeader(404)
					return
				}
				err = os.MkdirAll(filepath.Dir(absPath), 0755)
				if err != nil {
					log.Println(err)
					w.WriteHeader(500)
					return
				}
				file, err := os.OpenFile(absPath, os.O_CREATE|os.O_WRONLY, 0666)
				if err != nil {
					log.Println(err)
					w.WriteHeader(500)
					return
				}
				defer file.Close()
				_, err = io.Copy(file, resp.Body)
				if err != nil {
					log.Println(err)
					w.WriteHeader(500)
					return
				}
			} else {
				w.WriteHeader(404)
				return
			}
		}
		f, err := storage.Get(absPath)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		defer f.Close()
		ext := filepath.Ext(absPath)
		if len(ext) > 0 {
			ext = ext[1:]
		}
		if mime, ok := config.C.Mime[ext]; ok {
			w.Header().Set("Content-Type", mime)
		}
		io.Copy(w, f)
	})
	return router
}
