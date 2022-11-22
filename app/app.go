package app

import (
	"todolist/api"
	taskrepo "todolist/repository/task"
	userrepo "todolist/repository/user"
	"todolist/service/task"
	"todolist/service/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	echo *echo.Echo
	api  *api.Api
}

func New() *App {
	userService := user.NewService(userrepo.NewMapRepository())
	taskService := task.NewService(taskrepo.NewMapRepository())

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	return &App{
		echo: e,
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
	a.echo.DELETE("/user", a.api.HandleDeleteUser)
	a.echo.PUT("/task", a.api.HandleDoneTask)
}
