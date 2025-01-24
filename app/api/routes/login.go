package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/api/handlers"
	"github.com/golang-documented-todo-api/app/datasources/db"
)

func LoginRouter(api fiber.Router, service db.Database) {
	loginG := api.Group("/login")

	loginG.Get("/github", handlers.GetGithubRoute())
	loginG.Get("/github/callback", handlers.GetGithubCallbackRoute(service))
	loginG.Get("/google", handlers.GetGoogleRoute())
	loginG.Get("/google/callback", handlers.GetGoogleCallbackRoute(service))
}
