package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/techhubies/lenslocked/internal/handlers"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handlers.Get)

	return r
}
