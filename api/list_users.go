package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleListUsers(c echo.Context) error {
	users := a.userService.List()

	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	result := string(bytes)

	return c.JSON(http.StatusOK, result)
}
