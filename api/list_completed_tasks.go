package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleListCompletedTasks(c echo.Context) error {
	tasks := a.taskService.ListCompletedTasks()

	return c.JSON(http.StatusOK, tasks)
}
