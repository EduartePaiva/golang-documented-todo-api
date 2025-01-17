package handlers

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

func GetOpenApiSpec() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		_, b, _, _ := runtime.Caller(0)
		basepath := filepath.Dir(b)
		fmt.Println("---------- log here:", basepath)
		return c.SendFile("./static/docs/openapi.yaml", true)
	}
}

func GetScalarHtml() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendFile("./static/docs/scalar.html", true)
	}
}
