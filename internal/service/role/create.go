package role

import (
	"context"
	"task_manager/internal/dto"
)

func (s *Service) Create(ctx context.Context, role *dto.Role) error {
	err := s.roleRepo.CreateRole(ctx, role)
	if err != nil {
		return err
	}

	return nil
}
