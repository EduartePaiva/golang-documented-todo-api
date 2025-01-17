package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/api/routes"
)

func CreateApp() *fiber.App {
	app := fiber.New()
	api := app.Group("/api/v1")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})
	routes.DocsRouter(api)
	return app
}
