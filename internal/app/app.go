package app

import (
	"log"

	"github.com/techhubies/lenslocked/internal/controllers"
	"github.com/techhubies/lenslocked/internal/models"
)

// Application holds the dependencies for the web application.
type Application struct {
	Logger      *log.Logger
	UserStore   *models.UserStore
	UserHandler *controllers.UserHandler
}
