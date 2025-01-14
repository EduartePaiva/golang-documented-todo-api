package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/api/docs"
)

func CreateApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	app.Mount("/docs", docs.GetDocsApi())

	return app
}
