package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/datasources/db"
	"github.com/golang-documented-todo-api/app/pkg/arctic"
	"github.com/golang-documented-todo-api/app/pkg/env"
	"github.com/golang-documented-todo-api/app/pkg/session"
	"github.com/golang-documented-todo-api/app/repository"
	"github.com/jackc/pgx/v5/pgtype"
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
		userData := arctic.GithubUserData{}
		dec := json.NewDecoder(response.Body)
		err = dec.Decode(&userData)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}

		newUser, err := db.GetOrCreateNewUserAndReturn(nil, c.Context(), repository.User{
			Username: userData.Name,
			AvatarUrl: pgtype.Text{
				String: userData.AvatarURL,
				Valid:  true,
			},
			ProviderUserID: strconv.Itoa(int(userData.ID)),
			ProviderName:   repository.ProviderNameGithub,
		})
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}
		sessionToken, err := session.GenerateSessionToken()
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}
		session := session.CreateSession(c.Context(), nil, sessionToken, newUser.ID)

		return c.SendString("/github/callback")
	}
}
