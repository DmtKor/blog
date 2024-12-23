package main

import (
	"blog/src/commander"
	"blog/src/handlers"
	"blog/src/middleware"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.LoggerMiddleware, middleware.RedirectMiddleware)
	router.Get("/about", handlers.AboutHandler)
	router.Get("/", handlers.AboutHandler)
	router.Get("/browse", handlers.BrowseHandler)
	router.Get("/doc", handlers.SeePostHandler)
	router.Delete("/comment", handlers.DelCommHandler)
	router.Delete("/doc", handlers.DelPostHandler)
	router.Post("/comment", handlers.NewCommHandler)
	router.Post("/doc", handlers.NewPostHandler)
	commander.Comm.InitServer(initstr_temp, ":50050", router)
}