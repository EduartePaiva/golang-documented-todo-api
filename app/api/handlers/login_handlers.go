package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/pkg/arctic"
)

func GetGithubRoute() fiber.Handler {
	return func(c *fiber.Ctx) error {
		state, err := arctic.GenerateState()
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		return c.SendString("/github")
	}
}
func GetGithubCallbackRoute() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("/github/callback")
	}
}
