package board_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetBoardsBoardId(ctx echo.Context, boardId api.BoardId) error {
	board, err := h.boardService.GetBoardByID(ctx.Request().Context(), int64(boardId))
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusOK, board)
}
