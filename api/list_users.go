package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) HandleListUsers(c echo.Context) error {
	users := a.userService.List()

	return c.JSON(http.StatusOK, users)
}
