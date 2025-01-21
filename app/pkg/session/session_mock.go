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
