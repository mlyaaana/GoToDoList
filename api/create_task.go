package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"todolist/domain"
)

func (a *Api) HandleCreateTask(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")

	task := domain.NewTask(name, description)
	a.taskService.Create(task)

	return c.String(
		http.StatusCreated,
		fmt.Sprintf("Task created successfully (id: %s, name: %s, description: %s)", task.Id, name, description),
	)
}
