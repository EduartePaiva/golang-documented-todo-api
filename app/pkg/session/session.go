package session

import (
	"context"
	"crypto/rand"
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
