package session

import (
	"context"
	"testing"
	"time"

	"github.com/golang-documented-todo-api/app/pkg/crypto"
	"github.com/golang-documented-todo-api/app/pkg/encoding"
	"github.com/golang-documented-todo-api/app/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGenerateSessionToken(t *testing.T) {
	v, err := GenerateSessionToken()
	// This should never error
	t.Log("\nGenerated session token: ", v, "\n")
	assert.Nil(t, err)
	assert.NotEmpty(t, v)
	assert.Equal(t, len(v), 32)
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

func TestValidateSessionToken(t *testing.T) {
	// 1: select user, it it don't find it it should return an error
	testObj := new(SessionServiceMock)
	ctx := context.Background()
	// Simulates that the call to this function return no rows
	mockCall := testObj.On(
		"SelectUserBySessionID",
		ctx, encoding.EncodeHexLowerCase(crypto.Sha256([]byte("testing"))),
	).Return(repository.SelectUserBySessionIDRow{}, pgx.ErrNoRows)
	_, err := ValidateSessionToken(ctx, testObj, "testing")
	testObj.AssertExpectations(t)
	assert.ErrorIs(t, err, pgx.ErrNoRows)
	mockCall.Unset()

	// 2: if the session is expired, it should return an error with expired session
	mockCall = testObj.On(
		"SelectUserBySessionID",
		ctx, encoding.EncodeHexLowerCase(crypto.Sha256([]byte("testing"))),
	).Return(repository.SelectUserBySessionIDRow{
		ExpiresAt: pgtype.Timestamptz{Time: time.Date(2021, 0, 0, 0, 0, 0, 0, time.Local)},
	}, nil)
	_, err = ValidateSessionToken(ctx, testObj, "testing")
	assert.EqualError(t, err, "the token expired")
	mockCall.Unset()

	// 3: if the session is at least 15 days old it'll renew it, call the database and return the data
	mockCall = testObj.On(
		"SelectUserBySessionID",
		ctx, encoding.EncodeHexLowerCase(crypto.Sha256([]byte("testing"))),
	).Return(repository.SelectUserBySessionIDRow{
		ExpiresAt: pgtype.Timestamptz{Time: time.Now().Add(time.Hour * 24 * 14), Valid: true},
		ID_2:      "testing",
	}, nil).On(
		"UpdateSessionExpiresAt",
		mock.Anything,
		mock.Anything,
	).Return(nil)

	value, err := ValidateSessionToken(ctx, testObj, "testing")
	assert.NoError(t, err)
	// assert that the expiresAt is at least 29 days from now
	assert.Equal(t, value.ExpiresAt.Time.Compare(time.Now().Add(time.Hour*24*29)), 1)
	mockCall.Unset()

}

func TestValidateSessionToken2(t *testing.T) {
	testObj := new(SessionServiceMock)
	ctx := context.Background()
	// 4: if the expiredAt will end in more than 15 days it'll just return the data

	sessionData := repository.SelectUserBySessionIDRow{
		ExpiresAt: pgtype.Timestamptz{
			Time: time.Now().Add(time.Hour*24*15 + time.Hour),
		},
		Username: "hello test",
	}

	testObj.On(
		"SelectUserBySessionID",
		ctx, encoding.EncodeHexLowerCase(crypto.Sha256([]byte("testing"))),
	).Return(sessionData, nil)
	data, err := ValidateSessionToken(ctx, testObj, "testing")
	testObj.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, sessionData, data)
}
