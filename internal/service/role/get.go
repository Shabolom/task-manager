package role

import (
	"context"
	"task_manager/internal/dto"

	"go.uber.org/zap"
)

func (s *Service) Get(ctx context.Context, roleID int64) (*dto.Role, error) {
	id, err := s.roleRepo.GetRoleByID(ctx, roleID)
	if err != nil {
		s.logger.Warn("Failed to get role", zap.Error(err))
		return nil, err
	}

	return id, nil
}
