package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleDoneTask(c echo.Context) error {
	id := c.FormValue("id")

	a.taskService.Done(id)

	return c.NoContent(http.StatusOK)
}
