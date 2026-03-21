package task_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetBoardsBoardIdTasks(
	ctx echo.Context,
	boardId api.BoardId,
	params api.GetBoardsBoardIdTasksParams,
) error {

	if params.StatusId != nil {
		tasks, err := h.taskService.ListTasksByStatus(ctx.Request().Context(), int64(*params.StatusId))
		if err != nil {
			return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
		}
		return render.JSON(ctx, http.StatusOK, tasks)
	}

	tasks, err := h.taskService.ListTasksByBoard(ctx.Request().Context(), int64(boardId))
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusOK, tasks)
}
