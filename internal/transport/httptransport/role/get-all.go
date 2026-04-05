package role

import (
	"fmt"
	"net/http"
	"strconv"
	"task_manager/internal/dto"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetRoles(ctx echo.Context) error {
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

	roles, err := h.roleService.GetAll(ctx.Request().Context(), pagination)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return ctx.JSON(http.StatusOK, roles)
}
