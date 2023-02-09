package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) HandleListTasks(c echo.Context) error {
	tasks, err := a.taskService.List()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, NewTasks(tasks))
}
