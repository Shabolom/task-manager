package board_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/dto"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PatchBoardsBoardId(ctx echo.Context, boardId api.BoardId) error {
	request := new(dto.UpdateBoardRequest)
	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	if ctx.Echo().Validator != nil {
		if err := ctx.Validate(request); err != nil {
			return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
		}
	}

	updatedAt, err := h.boardService.UpdateBoard(ctx.Request().Context(), int64(boardId), request.Title, request.Description)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusOK, map[string]interface{}{
		"board_id":   boardId,
		"updated_at": updatedAt,
	})
}
