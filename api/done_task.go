package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) HandleDoneTask(c echo.Context) error {
	id := c.FormValue("id")

	a.taskService.Done(id)

	return c.NoContent(http.StatusOK)
}
