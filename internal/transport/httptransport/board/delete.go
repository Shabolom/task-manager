package board_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteBoardsBoardId(ctx echo.Context, boardId api.BoardId) error {
	deleted, err := h.boardService.DeleteBoard(ctx.Request().Context(), int64(boardId))
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	if !deleted {
		return render.BadRequest(ctx, render.CodeBadRequest, "Board not found", nil)
	}

	return render.JSON(ctx, http.StatusOK, map[string]string{
		"message": "Board deleted successfully",
	})
}
