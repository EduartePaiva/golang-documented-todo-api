package docs

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetDocsApi() *fiber.App {
	app := fiber.New()

	app.Get("/openapi.yaml", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		fmt.Println("hello world")
		fmt.Println("hello world")
		return c.SendFile("./api/docs/openapi.yaml", true)
	})

	app.Get("/redoc", func(c *fiber.Ctx) error {
		return c.SendFile("./api/docs/redoc.html", true)
	})

	return app
}
