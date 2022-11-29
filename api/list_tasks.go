package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) HandleListTasks(c echo.Context) error {
	tasks := a.taskService.List()

	return c.JSON(http.StatusOK, NewTasks(tasks))
}
