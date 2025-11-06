package router

import (
	"net/http"
	"small_imgbed/handler"
	"small_imgbed/middleware"
)

func SetUpRouter() http.Handler {
	r := http.NewServeMux()
	r.Handle("POST /upload", middleware.Auth(handler.Upload()))
	r.Handle("GET /", handler.Get())
	return r
}
