package handlers

import (
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/pkg/env"
)

func GetOpenApiSpec() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.SendFile(path.Join(env.Get().BasePath, "/app/static/docs/openapi.yaml"), true)
	}
}

func GetScalarHtml() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendFile(path.Join(env.Get().BasePath, "/app/static/docs/scalar.html"), true)
	}
}
