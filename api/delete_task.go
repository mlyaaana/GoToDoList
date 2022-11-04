package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleDeleteTask(c echo.Context) error {
	id := c.FormValue("id")

	name := a.taskService.Get(id).Name
	a.taskService.Delete(id)

	return c.String(http.StatusOK, fmt.Sprintf("Task deleted successfully (name: %s)", name))
}
