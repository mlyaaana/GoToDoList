package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleGetTask(c echo.Context) error {
	id := c.FormValue("id")

	task := a.taskService.Get(id)

	return c.JSON(http.StatusOK, NewTask(task))
}
