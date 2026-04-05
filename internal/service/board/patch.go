package board

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
)

func (s *Service) UpdateBoard(ctx context.Context, boardID int64, title, description string) (*time.Time, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	updatedAt, err := s.boardRepo.UpdateBoard(ctx, boardID, title, description)
	if err != nil {
		s.logger.Warn("Failed to updateBoard", zap.Error(err))
		return nil, err
	}
	if updatedAt == nil {
		return nil, errors.New("board not found")
	}
	return updatedAt, nil
}
