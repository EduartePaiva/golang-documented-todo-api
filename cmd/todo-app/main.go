package main

import (
	"log"

	"github.com/golang-documented-todo-api/api"
)

func main() {
	app := api.CreateApp()

	log.Fatal(app.Listen(":3000"))
}
