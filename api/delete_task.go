package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) HandleDeleteTask(c echo.Context) error {
	id := c.FormValue("id")

	a.taskService.Delete(id)

	return c.NoContent(http.StatusOK)
}
