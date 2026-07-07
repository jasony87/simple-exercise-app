package handlers

import (
	"errors"
	"net/http"
	"time"

	xapi "github.com/jasony87/simple-exercise-app/generated"
	store "github.com/jasony87/simple-exercise-app/internal/store"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	store *store.MemoryStore
}

func NewHandler(s *store.MemoryStore) *Handler {
	return &Handler{
		store: s,
	}
}

func (h *Handler) GetApiExerciseLogs(ctx echo.Context, params xapi.GetApiExerciseLogsParams) error {
	logItems, err := h.store.GetByDate(params.Date.Format(time.DateOnly))
	if err != nil {
		return handleError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, xapi.LogListResponse{
		Logs: &logItems,
	})
}

func (h *Handler) PutApiExerciseLogs(ctx echo.Context, params xapi.PutApiExerciseLogsParams) error {
	var logItem xapi.LogItem
	err := ctx.Bind(&logItem)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, xapi.Error{
			Code:    ptr(http.StatusBadRequest),
			Message: ptr(err.Error()),
		})
	}
	err = h.store.Update(params.Date.Format(time.DateOnly), logItem)
	if err != nil {
		return handleError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, logItem)
}

func (h *Handler) PostApiExerciseLogs(ctx echo.Context, params xapi.PostApiExerciseLogsParams) error {
	var logItem xapi.LogItem
	err := ctx.Bind(&logItem)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, xapi.Error{
			Code:    ptr(http.StatusBadRequest),
			Message: ptr(err.Error()),
		})
	}
	err = h.store.Add(params.Date.Format(time.DateOnly), logItem)
	if err != nil {
		return handleError(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, logItem)
}

func (h *Handler) GetApiFoodLogs(ctx echo.Context, params xapi.GetApiFoodLogsParams) error {
	logItems, err := h.store.GetByDate(params.Date.Format(time.DateOnly))
	if err != nil {
		return handleError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, xapi.LogListResponse{
		Logs: &logItems,
	})
}

// Helper function for pointer values
func ptr[T any](v T) *T {
	return &v
}

func handleError(ctx echo.Context, err error) error {
	if errors.Is(err, store.ErrLogEntryNotFound) {
		return ctx.JSON(http.StatusNotFound, xapi.Error{
			Code:    ptr(http.StatusNotFound),
			Message: ptr(err.Error()),
		})
	}
	if errors.Is(err, store.ErrLogEntryIdRequired) {
		return ctx.JSON(http.StatusBadRequest, xapi.Error{
			Code:    ptr(http.StatusBadRequest),
			Message: ptr(err.Error()),
		})
	}

	return ctx.JSON(http.StatusInternalServerError, xapi.Error{
		Code:    ptr(http.StatusInternalServerError),
		Message: ptr(err.Error()),
	})
}
