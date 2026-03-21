package status_service

import (
	"context"
	"errors"
	"task_manager/internal/dto"
	"time"
)

func (s *Service) CreateStatus(ctx context.Context, status *dto.Status) error {
	if status.Name == "" {
		return errors.New("name cannot be empty")
	}
	if status.Color == "" {
		status.Color = "#000000"
	}
	if status.BoardID == 0 {
		return errors.New("boardID cannot be zero")
	}

	status.CreatedAt = time.Now()
	status.UpdatedAt = time.Now()

	return s.statusRepo.CreateStatus(ctx, status)
}
