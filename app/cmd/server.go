package cmd

import "github.com/golang-documented-todo-api/app/api"

func RunServer() error {
	app := api.CreateApp()
	return app.Listen(":3000")
}
