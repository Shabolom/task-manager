package board_handler

import (
	"net/http"
	"task_manager/internal/dto"
	"task_manager/internal/render"
	"task_manager/pkg/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PostBoards(ctx echo.Context) error {
	request := new(dto.Board)

	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	if ctx.Echo().Validator != nil {
		if err := ctx.Validate(request); err != nil {
			return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
		}
	}

	userID, err := utils.ExtractUserIDFromToken(
		ctx.Request().Header.Get("Authorization"),
		h.jwtSecret,
	)
	if err != nil {
		return render.Unauthorized(ctx, err)
	}

	request.OwnerID = userID

	if err := h.boardService.CreateBoard(ctx.Request().Context(), request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusCreated, request)
}
