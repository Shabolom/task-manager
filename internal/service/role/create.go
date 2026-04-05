package role

import (
	"context"
	"task_manager/internal/dto"

	"go.uber.org/zap"
)

func (s *Service) Create(ctx context.Context, role *dto.Role) error {
	err := s.roleRepo.CreateRole(ctx, role)
	if err != nil {
		s.logger.Warn("Failed to CreateRole", zap.Error(err))
		return err
	}

	return nil
}
