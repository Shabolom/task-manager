package status_service

import "context"

func (s *Service) DeleteStatus(ctx context.Context, statusID int64) (bool, error) {
	rowsAffected, err := s.statusRepo.DeleteStatus(ctx, statusID)
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}
