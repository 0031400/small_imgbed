package router

import (
	"net/http"
	"small_imgbed/handler"
	"small_imgbed/middleware"
)

func SetUpRouter() http.Handler {
	r := http.NewServeMux()
	r.Handle("/login", middleware.Cors(middleware.Auth(handler.Login())))
	r.Handle("/upload", middleware.Cors(middleware.Auth(handler.Upload())))
	r.Handle("/", handler.Get())
	return middleware.Recover(r)
}
