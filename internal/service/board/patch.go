package board

import (
	"context"
	"errors"
	"time"
)

func (s *Service) UpdateBoard(ctx context.Context, boardID int64, title, description string) (*time.Time, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	updatedAt, err := s.boardRepo.UpdateBoard(ctx, boardID, title, description)
	if err != nil {
		return nil, err
	}
	if updatedAt == nil {
		return nil, errors.New("board not found")
	}
	return updatedAt, nil
}
