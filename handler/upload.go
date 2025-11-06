package handler

import (
	"log"
	"net/http"
	"small_imgbed/internal/rename"
	"small_imgbed/internal/storage"
)

func Upload() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, h, err := r.FormFile("file")
		if err != nil {
			log.Panicln(err)
		}
		newPath := rename.RenamePath(h.Filename)
		storage.Save(newPath, f)
		w.Write([]byte(newPath))
	})
	return router
}
