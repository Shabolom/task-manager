package di

import (
	auth_service "task_manager/internal/service/auth"
	board_service "task_manager/internal/service/board"
	role_service "task_manager/internal/service/role"
	status_service "task_manager/internal/service/status"
	task_service "task_manager/internal/service/task"
	user_service "task_manager/internal/service/user"
)

func (d *DI) GetRoleService() *role_service.Service {
	return role_service.New(d.GetRoleRepo())
}

func (d *DI) GetBoardService() *board_service.Service {
	return board_service.New(d.GetBoardRepo())
}

func (d *DI) GetUserService() *user_service.Service {
	return user_service.New(d.GetUserRepo())
}

func (d *DI) GetStatusService() *status_service.Service {
	return status_service.New(d.GetStatusRepo())
}

func (d *DI) GetTaskService() *task_service.Service {
	return task_service.New(d.GetTaskRepo())
}

func (d *DI) GetAuthService() *auth_service.Service {
	return auth_service.New(d.GetUserRepo(), "secret")
}
