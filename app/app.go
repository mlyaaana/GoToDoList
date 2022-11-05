package app

import (
	"github.com/labstack/echo/v4"
	"todolist/api"
	taskrepo "todolist/repository/task"
	userrepo "todolist/repository/user"
	"todolist/service/task"
	"todolist/service/user"
)

type App struct {
	echo *echo.Echo
	api  *api.Api
}

func New() *App {
	userService := user.NewService(userrepo.NewMapRepository())
	taskService := task.NewService(taskrepo.NewMapRepository())

	return &App{
		echo: echo.New(),
		api:  api.NewApi(userService, taskService),
	}
}

func (a *App) Start() {
	a.echo.Logger.Fatal(a.echo.Start(":8000"))
}

func (a *App) InitRoutes() {
	a.echo.GET("/tasks", a.api.HandleListTasks)
	a.echo.POST("/task", a.api.HandleCreateTask)
	a.echo.DELETE("/task", a.api.HandleDeleteTask)
	a.echo.GET("/task", a.api.HandleGetTask)
	a.echo.POST("/user", a.api.HandleCreateUser)
	a.echo.GET("/user", a.api.HandleGetUser)
	a.echo.GET("/users", a.api.HandleListUsers)
}
