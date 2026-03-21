package auth_handler

import (
	"net/http"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h Handler) PostAuthLogout(ctx echo.Context) error {
	return render.JSON(ctx, http.StatusOK, map[string]string{
		"message": "logged out",
	})
}
