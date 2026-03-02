package handler

import (
	"log"
	"net/http"
	"small_imgbed/internal/rename"
	"small_imgbed/internal/storage"
)

func Upload() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		f, h, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}
		newPath, err := rename.RenamePath(h.Filename)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
		_, err = storage.Save(newPath, f)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
		w.Write([]byte(newPath))
	})
	return router
}
