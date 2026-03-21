package httptransport

import (
	api "task_manager/gen"
	auth_handler "task_manager/internal/transport/httptransport/auth"
	board_handler "task_manager/internal/transport/httptransport/board"
	"task_manager/internal/transport/httptransport/role"
	status_handler "task_manager/internal/transport/httptransport/status"
	task_handler "task_manager/internal/transport/httptransport/task"
	user_handler "task_manager/internal/transport/httptransport/user"

	"github.com/labstack/echo/v4"
)

type (
	GetRoleHandler   = role.Handler
	GetBoardHandler  = board_handler.Handler
	GetUsersHandler  = user_handler.Handler
	GetStatusHandler = status_handler.Handler
	GetTaskHandler   = task_handler.Handler
	GetAuthHandler   = auth_handler.Handler
)

type Handlers struct {
	*GetRoleHandler
	*GetBoardHandler
	*GetUsersHandler
	*GetStatusHandler
	*GetTaskHandler
	*GetAuthHandler
}

func New(
	getRoleHandler *GetRoleHandler,
	boardHandler *GetBoardHandler,
	usersHandler *GetUsersHandler,
	statusHandler *GetStatusHandler,
	taskHandler *GetTaskHandler,
	authHandler *GetAuthHandler,
) *Handlers {
	return &Handlers{
		GetRoleHandler:   getRoleHandler,
		GetBoardHandler:  boardHandler,
		GetUsersHandler:  usersHandler,
		GetStatusHandler: statusHandler,
		GetTaskHandler:   taskHandler,
		GetAuthHandler:   authHandler,
	}
}

func (h Handlers) GetTasksTaskIdHistory(ctx echo.Context, taskId api.TaskId) error {
	//TODO implement me
	panic("implement me")
}
