package main

import (
	"errors"
	"log"
	"net/http"

	xapi "github.com/jasony87/simple-exercise-app/generated"
	"github.com/jasony87/simple-exercise-app/internal/handlers"
	"github.com/jasony87/simple-exercise-app/internal/store"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// Instantiate your handler
	s := store.NewMemoryStore()
	apiHandler := handlers.NewHandler(s)

	// Register routes using the generated function
	xapi.RegisterHandlers(e, apiHandler)

	log.Println("Server running on :8080")
	if err := e.Start(":8080"); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Shutting down the server:", err)
	}
}
