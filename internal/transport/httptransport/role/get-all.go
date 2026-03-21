package role

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetRoles(ctx echo.Context) error {
	roles, err := h.roleService.GetAll(ctx.Request().Context())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return ctx.JSON(http.StatusOK, roles)
}
