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
			Name:     "github_oauth_state",
			Value:    state,
			Path:     "/",
			Secure:   env.Get().GoEnv == "production",
			HTTPOnly: true,
			MaxAge:   60 * 10,
			SameSite: "lax",
		})
		return c.Redirect(url)
	}
}
func GetGithubCallbackRoute(service db.Database) fiber.Handler {
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

		newUser, err := db.GetOrCreateNewUserAndReturn(service, c.Context(), repository.User{
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
		newSession, err := session.CreateSession(c.Context(), service, sessionToken, newUser.ID)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}
		session.SetSessionTokenCookie(sessionToken, newSession.ExpiresAt.Time, c.Cookie)
		return c.Redirect(env.Get().FrontendURL + "/")
	}
}
func GetGoogleRoute() fiber.Handler {
	google := arctic.Google(
		env.Get().OAuth2.Google.ClientID,
		env.Get().OAuth2.Google.ClientSecret,
		env.Get().OAuth2.Google.RedirectURI,
	)
	return func(c *fiber.Ctx) error {
		state, err := arctic.GenerateState()
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		codeVerifier, err := arctic.GenerateCodeVerifier()
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		url := google.CreateAuthorizationURL(state, codeVerifier, []string{"openid", "profile"})
		c.Cookie(&fiber.Cookie{
			Name:     "google_oauth_state",
			Value:    state,
			Path:     "/",
			Secure:   env.Get().GoEnv == "production",
			HTTPOnly: true,
			MaxAge:   60 * 10,
			SameSite: "lax",
		})
		c.Cookie(&fiber.Cookie{
			Name:     "google_code_verifier",
			Value:    codeVerifier,
			Path:     "/",
			Secure:   env.Get().GoEnv == "production",
			HTTPOnly: true,
			MaxAge:   60 * 10,
			SameSite: "lax",
		})
		return c.Redirect(url)
	}
}
func GetGoogleCallbackRoute(service db.Database) fiber.Handler {
	google := arctic.Google(
		env.Get().OAuth2.Google.ClientID,
		env.Get().OAuth2.Google.ClientSecret,
		env.Get().OAuth2.Google.RedirectURI,
	)
	return func(c *fiber.Ctx) error {
		state := c.Query("state")
		code := c.Query("code")
		storedState := c.Cookies("google_oauth_state")
		codeVerifier := c.Cookies("google_code_verifier")

		if state == "" || code == "" || storedState == "" || codeVerifier == "" || storedState != state {
			return c.SendStatus(http.StatusBadRequest)
		}

		tokens, err := google.ValidateAuthorizationCode(c.Context(), code, codeVerifier)
		if err != nil {
			fmt.Println(err)
			fmt.Println("error 1")
			return c.SendStatus(http.StatusInternalServerError)
		}
		idToken, err := tokens.IdToken()
		if err != nil {
			fmt.Println("google should return an idToken. Check if the scopes are properly set")
			fmt.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}
		userData := arctic.GoogleUserData{}
		err = arctic.DecodeIdToken(idToken, &userData)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}
		newUser, err := db.GetOrCreateNewUserAndReturn(service, c.Context(), repository.User{
			Username:       userData.Name,
			AvatarUrl:      pgtype.Text{String: userData.AvatarURL, Valid: len(userData.AvatarURL) > 0},
			ProviderUserID: userData.ID,
			ProviderName:   repository.ProviderNameGoogle,
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
		newSession, err := session.CreateSession(c.Context(), service, sessionToken, newUser.ID)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}
		session.SetSessionTokenCookie(sessionToken, newSession.ExpiresAt.Time, c.Cookie)
		return c.Redirect(env.Get().FrontendURL + "/")
	}
}

type getUserResponse struct {
	Username  string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

func GetUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, ok := session.GetStoredSession(c)
		if !ok {
			fmt.Println("a stored session was not found")
			return c.SendStatus(http.StatusInternalServerError)
		}
		response := getUserResponse{
			Username:  session.Username,
			AvatarURL: session.AvatarUrl.String,
		}
		return c.JSON(response)
	}
}
