package user_service

import "context"

func (s *Service) DeleteUser(ctx context.Context, userID int64) (bool, error) {
	rowsAffected, err := s.userRepo.DeleteUser(ctx, userID)
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}
