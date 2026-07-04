package handlers

import (
	"net/http"
	"time"

	mealxapi "github.com/jasony87/simple-exercise-app/generated"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetApiExerciseLogs(ctx echo.Context, params mealxapi.GetApiExerciseLogsParams) error {
	logs := []mealxapi.LogItem{
		{
			Id:        ptr("ex123"),
			Type:      ptr("exercise"),
			Timestamp: ptr(time.Now()),
			Details: &map[string]interface{}{
				"exercise": "Running",
				"duration": 30,
			},
		},
	}

	return ctx.JSON(http.StatusOK, mealxapi.LogListResponse{
		Logs: &logs,
	})
}

func (h *Handler) PutApiExerciseLogs(ctx echo.Context, params mealxapi.PutApiExerciseLogsParams) error {
	logs := []mealxapi.LogItem{
		{
			Id:        ptr("abc"),
			Type:      ptr("exercise"),
			Timestamp: ptr(time.Now()),
			Details: &map[string]interface{}{
				"exercise": "Running",
				"duration": 50,
			},
		},
	}

	return ctx.JSON(http.StatusOK, mealxapi.LogListResponse{
		Logs: &logs,
	})
}

func (h *Handler) GetApiFoodLogs(ctx echo.Context, params mealxapi.GetApiFoodLogsParams) error {
	logs := []mealxapi.LogItem{
		{
			Id:        ptr("fd123"),
			Type:      ptr("food"),
			Timestamp: ptr(time.Now()),
			Details: &map[string]interface{}{
				"meal":     "Lunch",
				"calories": 600,
			},
		},
	}

	return ctx.JSON(http.StatusOK, mealxapi.LogListResponse{
		Logs: &logs,
	})
}

// Helper functions for pointer values

func ptr[T any](v T) *T {
	return &v
}
