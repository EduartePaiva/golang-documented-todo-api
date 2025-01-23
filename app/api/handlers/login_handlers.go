package handlers

import (
	"fmt"
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

		tokens, err := github.ValidateAuthorizationCode(c.Context(), code)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusInternalServerError)
		}
		accessToken, err := tokens.AccessToken()
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusInternalServerError)
		}

		request, err := http.NewRequestWithContext(c.Context(), "GET", "https://api.github.com/user", nil)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusInternalServerError)
		}
		request.Header.Add("Authorization", "Bearer "+accessToken)

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}
		// TODO: read the github necessary data

		return c.SendString("/github/callback")
	}
}
