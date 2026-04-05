package task_handler

import (
	"net/http"
	"strconv"
	api "task_manager/gen"
	"task_manager/internal/dto"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetBoardsBoardIdTasks(
	ctx echo.Context,
	boardId api.BoardId,
	params api.GetBoardsBoardIdTasksParams,
) error {
	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid limit")
	}
	offset, err := strconv.Atoi(ctx.QueryParam("offset"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid offset")
	}

	pagination := dto.Pagination{
		int64(limit),
		int64(offset),
	}

	if params.StatusId != nil {
		tasks, err := h.taskService.ListTasksByStatus(ctx.Request().Context(), *params.StatusId, pagination)
		if err != nil {
			return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
		}
		return render.JSON(ctx, http.StatusOK, tasks)
	}

	tasks, err := h.taskService.ListTasksByBoard(ctx.Request().Context(), boardId, pagination)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusOK, tasks)
}
