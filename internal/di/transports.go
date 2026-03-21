package di

import (
	"task_manager/internal/transport/httptransport"
	auth_handler "task_manager/internal/transport/httptransport/auth"
	board_handler "task_manager/internal/transport/httptransport/board"
	role_handler "task_manager/internal/transport/httptransport/role"
	status_handler "task_manager/internal/transport/httptransport/status"
	task_handler "task_manager/internal/transport/httptransport/task"
	user_handler "task_manager/internal/transport/httptransport/user"
)

// СОЗДАТЬ ВСЕ ХЭНДЛЕРЫ
func (d *DI) GetHTTPHandlers() *httptransport.Handlers {
	return httptransport.New(
		d.GetRoleHandler(),
		d.GetBoardHandler(),
		d.GetUserHandler(),
		d.GetStatusHandler(),
		d.GetTaskHandler(),
		d.GetAuthHandler(),
	)
}

func (d *DI) GetRoleHandler() *role_handler.Handler {
	return role_handler.New(d.GetRoleService())
}

func (d *DI) GetBoardHandler() *board_handler.Handler {
	return board_handler.New(d.GetBoardService(), d.Config().Secret)
}

func (d *DI) GetUserHandler() *user_handler.Handler {
	return user_handler.New(d.GetUserService(), d.Config().Secret)
}

func (d *DI) GetStatusHandler() *status_handler.Handler {
	return status_handler.New(d.GetStatusService())
}

func (d *DI) GetTaskHandler() *task_handler.Handler {
	return task_handler.New(d.GetTaskService(), d.Config().Secret)
}

func (d *DI) GetAuthHandler() *auth_handler.Handler {
	return auth_handler.New(d.GetAuthService())
}
