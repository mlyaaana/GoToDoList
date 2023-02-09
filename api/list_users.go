package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) HandleListUsers(c echo.Context) error {
	users, err := a.userService.List()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}
