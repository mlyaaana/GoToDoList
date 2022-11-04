package api

import (
	"todolist/service/task"
	"todolist/service/user"
)

type Api struct {
	userService *user.Service
	taskService *task.Service
}

func NewApi(userService *user.Service, taskService *task.Service) *Api {
	return &Api{
		userService: userService,
		taskService: taskService,
	}
}
