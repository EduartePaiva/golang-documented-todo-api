package session

import (
	"context"
	"testing"

	"github.com/golang-documented-todo-api/app/pkg/crypto"
	"github.com/golang-documented-todo-api/app/pkg/encoding"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestGenerateSessionToken(t *testing.T) {
	v, err := GenerateSessionToken()
	t.Log(len(v), v)
	// This should never error
	assert.Nil(t, err)
	assert.NotEmpty(t, v)
}

func TestCreateSession(t *testing.T) {
	mockSession := new(SessionServiceMock)
	session := CreateSession(
		context.Background(),
		mockSession,
		"testing",
		pgtype.UUID{
			Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 16, 15, 16},
			Valid: true,
		})

	// test that the uuid is the same
	assert.Equal(t, session.UserID, pgtype.UUID{
		Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 16, 15, 16},
		Valid: true,
	})

	// test that the encoding is working
	assert.Equal(t, session.ID, encoding.EncodeHexLowerCase(crypto.Sha256([]byte("testing"))))
}
