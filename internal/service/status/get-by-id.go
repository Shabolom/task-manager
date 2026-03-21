package status_service

import (
	"context"
	"errors"
	"task_manager/internal/dto"
)

func (s *Service) GetStatusByID(ctx context.Context, statusID int64) (*dto.Status, error) {
	status, err := s.statusRepo.GetStatusByID(ctx, statusID)
	if err != nil {
		return nil, err
	}
	if status == nil {
		return nil, errors.New("status not found")
	}
	return status, nil
}
