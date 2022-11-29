package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) HandleGetUser(c echo.Context) error {
	id := c.FormValue("id")

	user := a.userService.Get(id)

	return c.JSON(http.StatusOK, user)
}
