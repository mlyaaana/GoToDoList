package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleDeleteUser(c echo.Context) error {
	id := c.FormValue("id")

	name := a.userService.Get(id).Name
	a.userService.Delete(id)

	return c.String(http.StatusOK, fmt.Sprintf("User successfully deleted (name: %s)", name))
}
