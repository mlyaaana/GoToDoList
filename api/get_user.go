package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleGetUser(c echo.Context) error {
	id := c.FormValue("id")

	user := a.userService.Get(id)

	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	result := string(bytes)

	return c.JSON(http.StatusOK, result)
}
