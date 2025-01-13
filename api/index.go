package api

import "github.com/gofiber/fiber/v2"

func CreateApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	return app
}
