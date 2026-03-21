package status_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetBoardsBoardIdStatuses(ctx echo.Context, boardId api.BoardId) error {
	statuses, err := h.statusService.ListStatusesByBoard(ctx.Request().Context(), int64(boardId))
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}
	return render.JSON(ctx, http.StatusOK, statuses)
}
