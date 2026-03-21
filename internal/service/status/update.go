package status_service

import (
	"context"
	"errors"
	"time"
)

func (s *Service) UpdateStatus(ctx context.Context, statusID int64, name, color string) (*time.Time, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if color == "" {
		return nil, errors.New("color cannot be empty")
	}

	updatedAt, err := s.statusRepo.UpdateStatus(ctx, statusID, name, color)
	if err != nil {
		return nil, err
	}
	if updatedAt == nil {
		return nil, errors.New("status not found")
	}
	return updatedAt, nil
}
