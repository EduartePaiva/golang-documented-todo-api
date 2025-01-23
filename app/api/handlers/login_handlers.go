package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/pkg/arctic"
	"github.com/golang-documented-todo-api/app/pkg/env"
)

func GetGithubRoute() fiber.Handler {
	github := arctic.GitHub(
		env.Get().OAuth2.GitHub.ClientID,
		env.Get().OAuth2.GitHub.ClientSecret,
		env.Get().OAuth2.GitHub.RedirectURI,
	)
	return func(c *fiber.Ctx) error {
		state, err := arctic.GenerateState()
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		url := github.CreateAuthorizationURL(state, []string{})
		c.Cookie(&fiber.Cookie{
			Path:     "/",
			Secure:   env.Get().GoEnv == "production",
			HTTPOnly: true,
			MaxAge:   60 * 10,
			SameSite: "lax",
		})
		return c.Redirect(url)
	}
}
func GetGithubCallbackRoute() fiber.Handler {
	github := arctic.GitHub(
		env.Get().OAuth2.GitHub.ClientID,
		env.Get().OAuth2.GitHub.ClientSecret,
		env.Get().OAuth2.GitHub.RedirectURI,
	)
	return func(c *fiber.Ctx) error {
		state := c.Query("state")
		code := c.Query("code")
		storedState := c.Cookies("github_oauth_state")

		if state == "" || code == "" || storedState == "" || storedState != state {
			return c.SendStatus(http.StatusBadRequest)
		}

		tokens, err := github

		return c.SendString("/github/callback")
	}
}
