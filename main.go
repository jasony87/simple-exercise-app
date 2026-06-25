package main

import (
	"log"
	mealxapi "main/generated"
	"main/internal/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Instantiate your handler
	apiHandler := handlers.NewHandler()

	// Register routes using the generated function
	mealxapi.RegisterHandlers(e, apiHandler)

	log.Println("Server running on :8080")
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal("Shutting down the server:", err)
	}
}
