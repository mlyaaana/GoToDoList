package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) HandleDoneTask(c echo.Context) error {
	id := c.FormValue("id")

	err := a.taskService.Done(id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
