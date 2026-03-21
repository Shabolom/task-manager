package board

import (
	"context"
)

func (s *Service) DeleteBoard(ctx context.Context, boardID int64) (bool, error) {
	rowsAffected, err := s.boardRepo.DeleteBoard(ctx, boardID)
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}
