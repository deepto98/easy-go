package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *Application) routes() *chi.Mux {
	//Middleware

	//Routes
	a.App.Routes.Get("/", a.Handlers.Home)

	//static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
