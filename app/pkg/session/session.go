package session

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/golang-documented-todo-api/app/pkg/crypto"
	"github.com/golang-documented-todo-api/app/pkg/encoding"
	"github.com/golang-documented-todo-api/app/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

func GenerateSessionToken() (string, error) {
	// Generate a random slice of byte of length 20
	bytes := make([]byte, 0, 20)
	for i := 0; i < 20; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(1<<8))
		if err != nil {
			return "", err
		}
		bytes = append(bytes, byte(num.Uint64()))
	}
	token := encoding.EncodeBase32LowerCaseNoPadding(bytes)
	return token, nil
}

func CreateSession(
	ctx context.Context,
	service SessionService,
	token string,
	userId pgtype.UUID,
) repository.Session {
	sessionId := encoding.EncodeHexLowerCase(crypto.Sha256([]byte(token)))

	session := repository.CreateSessionParams{
		ID:     sessionId,
		UserID: userId,
		ExpiresAt: pgtype.Timestamptz{
			Time: time.Now().Add(time.Hour * 24 * 30),
		},
	}
	service.CreateSession(ctx, session)
	return repository.Session(session)
}

func ValidateSessionToken(ctx context.Context, service SessionService, token string) (repository.SelectUserBySessionIDRow, error) {
	sessionId := encoding.EncodeHexLowerCase(crypto.Sha256([]byte(token)))

	result, err := service.SelectUserBySessionID(ctx, sessionId)
	if err != nil {
		return result, err
	}

	if time.Now().Compare(result.ExpiresAt.Time) == 1 {
		return result, fmt.Errorf("the token expired")
	}

	if time.Now().Compare(result.ExpiresAt.Time.Add(time.Hour*24*15)) == 1 {
		// update expiredAt
	}

	return result, nil
}
