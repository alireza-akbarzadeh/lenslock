package app

import (
	"database/sql"
	"log"

	"github.com/techhubies/lenslocked/internal/handlers"
)

// Application holds the dependencies for the web application.
type Application struct {
	Logger  *log.Logger
	DB      *sql.DB
	handler *handlers.Handler

	// Add more fields if needed (e.g., Config, Cache, etc.)
}
