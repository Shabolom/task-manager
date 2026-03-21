package status_handler

import (
	"context"
	"task_manager/internal/dto"
	"time"
)

type StatusService interface {
	CreateStatus(ctx context.Context, status *dto.Status) error
	DeleteStatus(ctx context.Context, statusID int64) (bool, error)
	GetStatusByID(ctx context.Context, statusID int64) (*dto.Status, error)
	ListStatusesByBoard(ctx context.Context, boardID int64) ([]dto.Status, error)
	UpdateStatus(ctx context.Context, statusID int64, name, color string) (*time.Time, error)
}

type Handler struct {
	statusService StatusService
}

func New(statusRepo StatusService) *Handler {
	return &Handler{statusService: statusRepo}
}
