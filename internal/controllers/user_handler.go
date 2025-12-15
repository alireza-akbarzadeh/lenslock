package handlers

import (
	"net/http"

	"github.com/techhubies/lenslocked/internal/app"
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	app *app.Application
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(application *app.Application) *UserHandler {
	return &UserHandler{
		app: application,
	}
}

// Get Example handler method
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.app.Logger.Println("Home handler called")
	w.Write([]byte("Welcome to the Home Page!"))
}
