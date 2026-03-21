package status_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/dto"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PostBoardsBoardIdStatuses(ctx echo.Context, boardId api.BoardId) error {
	request := new(dto.Status)
	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	request.BoardID = int64(boardId)

	if ctx.Echo().Validator != nil {
		if err := ctx.Validate(request); err != nil {
			return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
		}
	}

	if err := h.statusService.CreateStatus(ctx.Request().Context(), request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusCreated, request)
}
