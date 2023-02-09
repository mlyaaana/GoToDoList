package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) HandleGetTask(c echo.Context) error {
	id := c.FormValue("id")

	task, err := a.taskService.Get(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, NewTask(task))
}
