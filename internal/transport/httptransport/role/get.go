package role

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetRolesRoleId(ctx echo.Context, roleId int64) error {
	role, err := h.roleService.Get(ctx.Request().Context(), roleId)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return ctx.JSON(http.StatusOK, role)
}
