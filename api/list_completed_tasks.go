package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) HandleListCompletedTasks(c echo.Context) error {
	tasks, err := a.taskService.ListCompletedTasks()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tasks)
}
