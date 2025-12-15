package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/techhubies/lenslocked/internal/app"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	app := app.New()
	r.Use(httprate.LimitByIP(app.Config.RateLimit, 1*time.Minute))

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/", HelloWorld)

	})

	//    // Private Routes
	// // Require Authentication
	// r.Group(func(r chi.Router) {
	//     r.Use(AuthMiddleware)
	//     r.Post("/manage", CreateAsset)
	// })

	http.ListenAndServe(":3000", r)
}
