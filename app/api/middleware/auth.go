package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/datasources/db"
	"github.com/golang-documented-todo-api/app/pkg/session"
)

// Authenticate the user and store the session for the next users
func AuthMiddleware(service db.SessionService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sessionCookie := c.Cookies("session")
		if sessionCookie == "" {
			return c.SendStatus(http.StatusUnauthorized)
		}
		result, err := session.ValidateSessionToken(c.Context(), service, sessionCookie)
		if err != nil {
			return c.SendStatus(http.StatusUnauthorized)
		}
		session.StoreSession(c, result)
		return c.Next()
	}
}
