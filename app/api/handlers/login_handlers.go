package handlers

import "github.com/gofiber/fiber/v2"

func GetGithubRoute() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("/github")
	}
}
func GetGithubCallbackRoute() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("/github/callback")
	}
}
