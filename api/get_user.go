package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleGetUser(c echo.Context) error {
	id := c.FormValue("id")

	user := a.userService.Get(id)

	return c.JSON(http.StatusOK, user)
}
