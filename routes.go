package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func routes(app *AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/home", Repo.index)
	mux.Get("/api/v1", Repo.zipfinder)
	return mux
}
