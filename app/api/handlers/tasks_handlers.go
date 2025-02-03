package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/datasources/db"
	"github.com/golang-documented-todo-api/app/pkg/session"
	"github.com/golang-documented-todo-api/app/pkg/tasks"
	"github.com/golang-documented-todo-api/app/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetTasks(service db.TasksServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userData, ok := session.GetStoredSession(c)
		if !ok {
			return c.SendStatus(http.StatusInternalServerError)
		}
		fmt.Println(userData)

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

		for i := 0; i < len(data); i++ {
			err = service.PostTask(c.Context(), repository.PostTaskParams{
				ID:          data[i].ID,
				UserID:      user.ID,
				TodoText:    data[i].Text,
				Done:        data[i].Done,
				UpdatedAt:   pgtype.Timestamp{Time: data[i].UpdatedAt, Valid: true},
				TodoText_2:  data[i].Text,
				Done_2:      data[i].Done,
				UpdatedAt_2: pgtype.Timestamp{Time: data[i].UpdatedAt, Valid: true},
			})
			if err != nil {
				fmt.Println("error while inserting at index:", i)
			}

		}

		return c.SendStatus(http.StatusAccepted)
	}
}
