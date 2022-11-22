package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"todolist/domain"
)

func (a *Api) HandleCreateTask(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")

	task := domain.NewTask(name, description)
	a.taskService.Create(task)

	return c.JSON(http.StatusCreated, NewTask(task))
}
