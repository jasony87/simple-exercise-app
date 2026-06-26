package handlers

import (
	mealxapi "main/generated"
	"net/http"
	"time"

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
			Timestamp: ptrTime(time.Now()),
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
			Timestamp: ptrTime(time.Now()),
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
			Timestamp: ptrTime(time.Now()),
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
func ptr(s string) *string {
	return &s
}

func ptrTime(t time.Time) *time.Time {
	return &t
}

// package handlers

// import (
// 	"context"
// 	mealxapi "main/generated"
// )

// type Handler struct {
// 	// you can add dependencies here, e.g. DB connection
// }

// func NewHandler() *Handler {
// 	return &Handler{}
// }

// func (h *Handler) GetApiExerciseLogs(ctx context.Context, params mealxapi.GetApiExerciseLogsParams) error {
// 	// GET api exercise logs
// 	return nil
// }

// // Get food logs by date
// // (GET /api/food/logs)
// func GetApiFoodLogs(ctx context.Context, params mealxapi.GetApiFoodLogsParams) error {
// 	return nil
// }
