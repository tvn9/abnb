package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tvn9/abnb/pkg/config"
	"github.com/tvn9/abnb/pkg/handlers"

	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fs := http.FileServer(http.Dir("./assets/"))
	mux.Handle("/assets/*", http.StripPrefix("/assets", fs))

	return mux
}
