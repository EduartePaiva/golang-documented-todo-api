package session

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/datasources/db"
	"github.com/golang-documented-todo-api/app/pkg/crypto"
	"github.com/golang-documented-todo-api/app/pkg/encoding"
	"github.com/golang-documented-todo-api/app/pkg/env"
	"github.com/golang-documented-todo-api/app/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

func GenerateSessionToken() (string, error) {
	// Generate a random slice of byte of length 20
	bytes, err := crypto.GetRandomValues(20)
	if err != nil {
		return "", err
	}
	token := encoding.EncodeBase32LowerCaseNoPadding(bytes)
	return token, nil
}

func CreateSession(
	ctx context.Context,
	service db.SessionService,
	token string,
	userId pgtype.UUID,
) (repository.Session, error) {
	sessionId := encoding.EncodeHexLowerCase(crypto.Sha256([]byte(token)))

	session := repository.CreateSessionParams{
		ID:     sessionId,
		UserID: userId,
		ExpiresAt: pgtype.Timestamptz{
			Time:  time.Now().Add(time.Hour * 24 * 30),
			Valid: true,
		},
	}
	err := service.CreateSession(ctx, session)
	return repository.Session(session), err
}

func ValidateSessionToken(
	ctx context.Context,
	service db.SessionService,
	token string,
) (repository.SelectUserBySessionIDRow, error) {
	sessionId := encoding.EncodeHexLowerCase(crypto.Sha256([]byte(token)))

	result, err := service.SelectUserBySessionID(ctx, sessionId)
	if err != nil {
		return result, err
	}

	if time.Now().Compare(result.ExpiresAt.Time) == 1 {
		return result, fmt.Errorf("the token expired")
	}

	if time.Now().Compare(result.ExpiresAt.Time.Add(-time.Hour*24*15)) == 1 {
		// update expiredAt
		result.ExpiresAt.Time = time.Now().Add(time.Hour * 24 * 30)
		service.UpdateSessionExpiresAt(
			ctx,
			repository.UpdateSessionExpiresAtParams{ExpiresAt: result.ExpiresAt, ID: result.ID_2},
		)
	}
	return result, nil
}

func SetSessionTokenCookie(
	sessionToken string,
	expiresAt time.Time,
	cookie func(cookie *fiber.Cookie),
) {
	cookie(&fiber.Cookie{
		Name:     "session",
		Value:    sessionToken,
		HTTPOnly: true,
		SameSite: "lax",
		Secure:   env.Get().GoEnv == "production",
		Expires:  expiresAt,
		Path:     "/",
	})
}

func DeleteSessionTokenCookie(cookie func(cookie *fiber.Cookie)) {
	cookie(&fiber.Cookie{
		Name:     "session",
		Value:    "",
		HTTPOnly: true,
		SameSite: "lax",
		Secure:   env.Get().GoEnv == "production",
		MaxAge:   0,
		Path:     "/",
	})
}

func InvalidateSession(ctx context.Context, service db.SessionService, sessionID string) error {
	return service.DeleteSessionByID(ctx, sessionID)
}
