package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleListTasks(c echo.Context) error {
	tasks := a.taskService.List()

	return c.JSON(http.StatusOK, NewTasks(tasks))
}
