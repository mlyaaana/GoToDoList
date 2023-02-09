package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"todolist/domain"
)

func (a *Api) HandleCreateTask(c echo.Context) error {
	body := struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Deadline    time.Time `json:"deadline"`
	}{}

	if err := c.Bind(&body); err != nil {
		return err
	}

	task := domain.NewTask(body.Name, body.Description, body.Deadline.UTC())
	err := a.taskService.Create(task)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, NewTask(task))
}
