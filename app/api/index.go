package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-documented-todo-api/app/api/routes"
	"github.com/golang-documented-todo-api/app/datasources"
	"github.com/golang-documented-todo-api/app/pkg/env"
)

func CreateApp(ctx context.Context, dataSource *datasources.DataSources) *fiber.App {
	app := fiber.New()
	api := app.Group("/api/v1")
	if env.Get().GoEnv != "production" {
		api.Use(logger.New())
	}

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})
	routes.DocsRouter(api)
	routes.LoginRouter(api, dataSource.DB)
	return app
}
