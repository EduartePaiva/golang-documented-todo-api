package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/api/handlers"
)

func DocsRouter(api fiber.Router) {
	api.Get("/docs/openapi.yaml", handlers.GetOpenApiSpec())
	api.Get("/docs/reference", handlers.GetScalarHtml())
}
