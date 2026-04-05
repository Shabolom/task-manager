package board_handler

import (
	"net/http"
	"strconv"
	"task_manager/internal/dto"
	"task_manager/internal/render"
	"task_manager/pkg/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetBoards(ctx echo.Context) error {
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

	//fmt.Println(ctx.Request().Header.Get("Authorization"))
	userID, err := utils.ExtractUserIDFromToken(
		ctx.Request().Header.Get("Authorization"),
		h.jwtSecret,
	)
	if err != nil {
		return render.Unauthorized(ctx, err)
	}

	boards, err := h.boardService.ListBoardsByOwner(ctx.Request().Context(), userID, pagination)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusOK, boards)
}
