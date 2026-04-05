package role

import (
	"context"
	"task_manager/internal/dto"

	"go.uber.org/zap"
)

func (s *Service) GetAll(ctx context.Context, pagination dto.Pagination) ([]dto.Role, error) {
	roles, err := s.roleRepo.ListRoles(ctx, pagination)
	if err != nil {
		s.logger.Warn("Failed to list roles", zap.Error(err))
		return nil, err
	}

	return roles, nil
}
