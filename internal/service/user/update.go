package user_service

import (
	"context"
	"errors"
	"strings"
	"task_manager/internal/dto"
	"time"
)

func (s *Service) UpdateUser(
	ctx context.Context,
	userID int64,
	request *dto.UpdateUserRequest,
) (*time.Time, error) {
	if request == nil {
		return nil, errors.New("request is nil")
	}

	if request.Email != nil {
		email := strings.TrimSpace(*request.Email)
		if email == "" {
			return nil, errors.New("email cannot be empty")
		}
		request.Email = &email
	}

	if request.FullName != nil {
		fullName := strings.TrimSpace(*request.FullName)
		if fullName == "" {
			return nil, errors.New("full_name cannot be empty")
		}
		request.FullName = &fullName
	}

	if request.RoleID != nil && *request.RoleID <= 0 {
		return nil, errors.New("role_id must be greater than 0")
	}

	updatedAt, err := s.userRepo.UpdateUser(ctx, userID, request)
	if err != nil {
		return nil, err
	}
	if updatedAt == nil {
		return nil, errors.New("user not found")
	}

	return updatedAt, nil
}
