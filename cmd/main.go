package main

import (
	"errors"
	"log"
	"net/http"

	mealxapi "github.com/jasony87/simple-exercise-app/generated"
	"github.com/jasony87/simple-exercise-app/internal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Instantiate your handler
	apiHandler := handlers.NewHandler()

	// Register routes using the generated function
	mealxapi.RegisterHandlers(e, apiHandler)

	log.Println("Server running on :8080")
	if err := e.Start(":8080"); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Shutting down the server:", err)
	}
}
