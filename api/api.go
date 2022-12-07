package api

import (
	"todolist/service/auth"
	"todolist/service/task"
	"todolist/service/user"
)

type Api struct {
	userService *user.Service
	taskService *task.Service
	authService *auth.Service
}

func NewApi(userService *user.Service, taskService *task.Service, authService *auth.Service) *Api {
	return &Api{
		userService: userService,
		taskService: taskService,
		authService: authService,
	}
}
