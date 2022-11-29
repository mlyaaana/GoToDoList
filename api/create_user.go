package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"todolist/domain"
)

func (a *Api) HandleCreateUser(c echo.Context) error {
	name := c.FormValue("name")

	user := domain.NewUser(name)
	a.userService.Create(user)

	return c.JSON(http.StatusCreated, user)
}
