package role

import (
	"context"
	"task_manager/internal/dto"

	"go.uber.org/zap"
)

func (s *Service) Patch(ctx context.Context, patchReq dto.UpdateRoleRequest) (dto.Role, error) {
	result, err := s.roleRepo.UpdateRole(ctx, patchReq)
	if err != nil {
		s.logger.Warn("Failed to updateRole", zap.Error(err))
		return dto.Role{}, err
	}

	return result, nil
}
