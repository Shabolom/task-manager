package role

import (
	"net/http"
	"task_manager/internal/dto"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PostRoles(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.Role)

	if err := c.Bind(requestData); err != nil {
		return render.BadRequest(c, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	if c.Echo().Validator != nil {
		if err := c.Validate(requestData); err != nil {
			return render.BadRequest(c, render.CodeBadRequest, render.MsgBadRequest, err)
		}
	}

	err := h.roleService.Create(ctx, requestData)
	if err != nil {
		return render.BadRequest(c, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(c, http.StatusOK, requestData)
}
