package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleAuth(c echo.Context) error {
	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&body); err != nil {
		return err
	}

	userId, err := a.authService.Auth(body.Username, body.Password)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, userId)
}
