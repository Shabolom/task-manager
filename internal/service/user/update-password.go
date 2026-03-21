package user_service

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) UpdateUserPassword(ctx context.Context, userID int64, passwordHash string) (bool, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwordHash), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}

	rowsAffected, err := s.userRepo.UpdateUserPassword(ctx, userID, hash)
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}
