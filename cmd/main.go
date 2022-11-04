package main

import "todolist/app"

func main() {
	application := app.New()
	application.InitRoutes()
	application.Start()
}
