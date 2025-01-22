package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/api/handlers"
)

func LoginRouter(api fiber.Router) {
	loginG := api.Group("/login")

	loginG.Get("/github", handlers.GetGithubRoute())
	loginG.Get("/github/callback", handlers.GetGithubCallbackRoute())
}
