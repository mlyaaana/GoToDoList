package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) HandleDeleteUser(c echo.Context) error {
	id := c.FormValue("id")

	a.userService.Delete(id)

	return c.NoContent(http.StatusOK)
}
