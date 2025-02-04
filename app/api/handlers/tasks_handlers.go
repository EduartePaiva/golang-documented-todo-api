package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/datasources/db"
	"github.com/golang-documented-todo-api/app/pkg/session"
	"github.com/golang-documented-todo-api/app/pkg/tasks"
)

func GetTasks(service db.TasksServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userData, ok := session.GetStoredSession(c)
		if !ok {
			return c.SendStatus(http.StatusInternalServerError)
		}
		tasks, err := tasks.GetTasksForUser(service, userData.ID, c.Context())
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusInternalServerError)
		}
		return c.JSON(tasks)
	}
}

func PostTasks(service db.TasksServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := session.GetStoredSession(c)
		if !ok {
			return c.SendStatus(http.StatusUnauthorized)
		}
		ctype := string(c.Request().Header.ContentType())
		if ctype != "application/json" {
			return c.SendStatus(http.StatusBadRequest)
		}
		data, err := tasks.ProcessAndValidateIncomingTasks(c.Body())
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}
		tasks.PostTasks(service, data, user.ID, c.Context())
		return c.SendString("tasks created")
	}
}
