package board_handler

import (
	"fmt"
	"net/http"
	"task_manager/internal/render"
	"task_manager/pkg/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetBoards(ctx echo.Context) error {
	fmt.Println(ctx.Request().Header.Get("Authorization"))
	userID, err := utils.ExtractUserIDFromToken(
		ctx.Request().Header.Get("Authorization"),
		h.jwtSecret,
	)
	if err != nil {
		return render.Unauthorized(ctx, err)
	}

	boards, err := h.boardService.ListBoardsByOwner(ctx.Request().Context(), userID)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusOK, boards)
}
