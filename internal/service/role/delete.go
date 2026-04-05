package role

import (
	"context"

	"go.uber.org/zap"
)

func (s *Service) DeleteRole(ctx context.Context, roleID int64) (int64, error) {
	res, err := s.roleRepo.DeleteRole(ctx, roleID)
	if err != nil {
		s.logger.Warn("Failed to deleteRole", zap.Error(err))
		return res, err
	}

	return res, nil
}
