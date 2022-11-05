package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"todolist/domain"
)

func (a *Api) HandleCreateUser(c echo.Context) error {
	name := c.FormValue("name")

	user := domain.NewUser(name)
	a.userService.Create(user)

	return c.String(http.StatusCreated, fmt.Sprintf("User successfully created (id: %s, name: %s)", user.Id, name))
}
