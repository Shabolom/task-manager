package status_service

import (
	"context"
	"task_manager/internal/dto"
	"time"
)

type StatusRepo interface {
	CreateStatus(ctx context.Context, status *dto.Status) error
	DeleteStatus(ctx context.Context, statusID int64) (int64, error)
	GetStatusByID(ctx context.Context, statusID int64) (*dto.Status, error)
	ListStatusesByBoard(ctx context.Context, boardID int64) ([]dto.Status, error)
	UpdateStatus(
		ctx context.Context,
		statusID int64,
		name string,
		color string,
	) (*time.Time, error)
}

type Service struct {
	statusRepo StatusRepo
}

func New(statusRepo StatusRepo) *Service {
	return &Service{statusRepo: statusRepo}
}
