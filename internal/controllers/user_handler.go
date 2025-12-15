package controllers

import (
	"net/http"

	"github.com/techhubies/lenslocked/internal/models"
	"github.com/techhubies/lenslocked/internal/views"
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	store *models.UserStore
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(store *models.UserStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

// Get Example handler method
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	// log.Println("User list handler called") // Optionally use the standard logger or inject a logger interface
	users := h.store.All()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := views.RenderUserList(w, users); err != nil {
		http.Error(w, "Failed to render user list", http.StatusInternalServerError)
	}
}
