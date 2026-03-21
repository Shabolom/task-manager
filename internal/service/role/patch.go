package role

import (
	"context"
	"task_manager/internal/dto"
)

func (s *Service) Patch(ctx context.Context, patchReq dto.UpdateRoleRequest) (dto.Role, error) {
	result, err := s.roleRepo.UpdateRole(ctx, patchReq)
	if err != nil {
		return dto.Role{}, err
	}

	return result, nil
}
