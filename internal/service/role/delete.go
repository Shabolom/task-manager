package role

import (
	"context"
)

func (s *Service) DeleteRole(ctx context.Context, roleID int64) (int64, error) {
	res, err := s.roleRepo.DeleteRole(ctx, roleID)
	if err != nil {
		return res, err
	}

	return res, nil
}
