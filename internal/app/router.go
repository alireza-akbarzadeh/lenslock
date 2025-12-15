package app

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func SetupRouter(app *Application, cfg *Config) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(cfg.RateLimit, time.Minute))
	r.Get("/", app.UserHandler.Get)
	return r
}
