package session

import (
	"context"

	"github.com/golang-documented-todo-api/app/repository"
	"github.com/stretchr/testify/mock"
)

type SessionServiceMock struct {
	mock.Mock
}

func (m *SessionServiceMock) CreateSession(ctx context.Context, arg repository.CreateSessionParams) error {
	return nil
}
func (m *SessionServiceMock) SelectUserBySessionID(ctx context.Context, id string) (repository.SelectUserBySessionIDRow, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(repository.SelectUserBySessionIDRow), args.Error(1)
}
func (m *SessionServiceMock) UpdateSessionExpiresAt(
	ctx context.Context,
	arg repository.UpdateSessionExpiresAtParams,
) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *SessionServiceMock) DeleteSessionByID(ctx context.Context, id string) error {
	return nil
}
