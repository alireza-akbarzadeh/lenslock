package main

import (
	"log"
	"net/http"
	"os"

	"github.com/techhubies/lenslocked/internal/app"
	"github.com/techhubies/lenslocked/internal/controllers"
	"github.com/techhubies/lenslocked/internal/models"
)

func main() {
	logger := log.New(os.Stdout, "[lenslocked] ", log.LstdFlags)
	cfg := app.DefaultConfig()
	application := &app.Application{
		Logger: logger,
	}
	userStore := models.NewUserStore()
	handler := controllers.NewUserHandler(application, userStore)

	r := app.SetupRouter(handler, cfg)
	logger.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Fatalf("Error starting server: %v", err)
	}
}
