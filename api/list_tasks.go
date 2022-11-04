package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleListTasks(c echo.Context) error {
	tasks := a.taskService.List()

	bytes, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	result := string(bytes)

	return c.JSON(http.StatusOK, result)
}
