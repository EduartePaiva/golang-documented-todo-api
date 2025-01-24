package db

import (
	"context"

	. "github.com/golang-documented-todo-api/app/repository"
)

// Service is an interface from which our api module can access our repository of all our models
type SessionService interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) error
	SelectUserBySessionID(ctx context.Context, id string) (SelectUserBySessionIDRow, error)
	UpdateSessionExpiresAt(ctx context.Context, arg UpdateSessionExpiresAtParams) error
}

type sessionService struct {
	db Database
}

func NewService(db Database) SessionService {
	return &sessionService{db: db}
}

func (s *sessionService) CreateSession(ctx context.Context, arg CreateSessionParams) error {
	return s.db.CreateSession(ctx, arg)
}

func (s *sessionService) SelectUserBySessionID(ctx context.Context, id string) (SelectUserBySessionIDRow, error) {
	return s.db.SelectUserBySessionID(ctx, id)
}

func (s *sessionService) UpdateSessionExpiresAt(ctx context.Context, arg UpdateSessionExpiresAtParams) error {
	return s.db.UpdateSessionExpiresAt(ctx, arg)
}

// Services that deals with the user table
type UserServices interface {
	SelectUserFromProviderNameAndId(
		ctx context.Context,
		arg SelectUserFromProviderNameAndIdParams,
	) (User, error)
	UpdateUserAvatarURL(ctx context.Context, arg UpdateUserAvatarURLParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
}
type userService struct {
	db Database
}

func (s *userService) SelectUserFromProviderNameAndId(
	ctx context.Context,
	arg SelectUserFromProviderNameAndIdParams,
) (User, error) {
	return s.db.SelectUserFromProviderNameAndId(ctx, arg)
}
