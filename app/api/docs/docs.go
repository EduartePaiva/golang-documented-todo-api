package docs

import (
	"github.com/gofiber/fiber/v2"
)

func GetDocsApi() *fiber.App {
	app := fiber.New()

	app.Get("/openapi.yaml", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.SendFile("./api/docs/openapi.yaml", true)
	})

	app.Get("/reference", func(c *fiber.Ctx) error {
		return c.SendFile("./api/docs/scalar.html", true)
	})

	return app
}
