package user_handler

import (
	"net/http"
	"strconv"
	"task_manager/internal/dto"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUsers(ctx echo.Context) error {
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

	users, err := h.userService.ListUsers(ctx.Request().Context(), pagination)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}
	return render.JSON(ctx, http.StatusOK, users)
}
