package main

import (
	"WebGo/pkg/config"
	"WebGo/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//	mux := pat.New()
	//
	//	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	//	return mux
	//

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//mux.Use(WriteToConsole)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
