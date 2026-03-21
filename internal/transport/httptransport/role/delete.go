package role

import (
	"net/http"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteRolesRoleId(c echo.Context, roleID int64) error {
	ctx := c.Request().Context()

	rows, err := h.roleService.DeleteRole(ctx, roleID)
	if err != nil {
		return render.BadRequest(c, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	if rows == 0 {
		return render.NotFound(c, render.CodeNotFound, render.MsgNotFound, nil)
	}

	return render.JSON(c, http.StatusOK, map[string]int64{
		"deleted": rows,
	})
}
