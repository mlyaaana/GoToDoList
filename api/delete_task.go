package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleDeleteTask(c echo.Context) error {
	id := c.FormValue("id")

	a.taskService.Delete(id)

	return c.NoContent(http.StatusOK)
}
