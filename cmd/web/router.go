package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hculpan/vinylbase/cmd/web/handlers"
)

func setRoutes(r *chi.Mux) {
	fileServer(r, "/static", http.Dir("./assets"))

	r.Group(func(r chi.Router) {
		r.Use(handlers.AuthMiddleware)

		// Define your protected routes here
		r.Get("/mycollection", handlers.MyCollectionHandler)
	})

	// Define unprotected routes here
	r.Get("/", handlers.RootHandler)
	r.Get("/login", handlers.LoginHandler)
	r.Post("/login", handlers.LoginPostHandler)
	r.Get("/error", handlers.ErrorHandler)
}
