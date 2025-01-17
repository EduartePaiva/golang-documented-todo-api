package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetOpenApiSpec() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.SendFile("./static/docs/openapi.yaml", true)
	}
}

func GetScalarHtml() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendFile("./static/docs/scalar.html", true)
	}
}
