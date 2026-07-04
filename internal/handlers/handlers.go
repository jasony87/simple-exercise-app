package handlers

import (
	"net/http"
	"time"

	xapi "github.com/jasony87/simple-exercise-app/generated"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetApiExerciseLogs(ctx echo.Context, params xapi.GetApiExerciseLogsParams) error {
	logs := []xapi.LogItem{
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

	return ctx.JSON(http.StatusOK, xapi.LogListResponse{
		Logs: &logs,
	})
}

func (h *Handler) PutApiExerciseLogs(ctx echo.Context, params xapi.PutApiExerciseLogsParams) error {
	logs := []xapi.LogItem{
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

	return ctx.JSON(http.StatusOK, xapi.LogListResponse{
		Logs: &logs,
	})
}

func (h *Handler) GetApiFoodLogs(ctx echo.Context, params xapi.GetApiFoodLogsParams) error {
	logs := []xapi.LogItem{
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

	return ctx.JSON(http.StatusOK, xapi.LogListResponse{
		Logs: &logs,
	})
}

// Helper functions for pointer values

func ptr[T any](v T) *T {
	return &v
}
