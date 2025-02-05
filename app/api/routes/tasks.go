package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/api/handlers"
	"github.com/golang-documented-todo-api/app/api/middleware"
	"github.com/golang-documented-todo-api/app/datasources/db"
)

func TaskRouter(api fiber.Router, service db.Database) {
	taskG := api.Group("/tasks", middleware.AuthMiddleware(service))

	taskG.Get("", handlers.GetTasks(service))
	taskG.Post("", handlers.PostTasks(service))
	taskG.Delete("/:id", handlers.DeleteTask(service))
}
