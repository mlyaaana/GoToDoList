package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) HandleListCompletedTasks(c echo.Context) error {
	tasks := a.taskService.ListCompletedTasks()

	return c.JSON(http.StatusOK, tasks)
}
