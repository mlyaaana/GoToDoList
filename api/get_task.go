package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleGetTask(c echo.Context) error {
	id := c.FormValue("id")

	task := a.taskService.Get(id)

	bytes, err := json.Marshal(task)
	if err != nil {
		return err
	}
	result := string(bytes)

	return c.JSON(http.StatusOK, result)
}