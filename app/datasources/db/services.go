package db

import (
	"context"

	"github.com/golang-documented-todo-api/app/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

/*
The main idea of those interfaces is to split up the database interface so mocking will be easier
and it'll guarantees that some functions don't have access to things that it don't need
*/

// Service is an interface from which our api module can access our repository of all our models
type SessionService interface {
	CreateSession(ctx context.Context, arg repository.CreateSessionParams) error
	SelectUserBySessionID(ctx context.Context, id string) (repository.SelectUserBySessionIDRow, error)
	UpdateSessionExpiresAt(ctx context.Context, arg repository.UpdateSessionExpiresAtParams) error
	DeleteSessionByID(ctx context.Context, id string) error
}

// Services that deals with the user table
type UserServices interface {
	SelectUserFromProviderNameAndId(
		ctx context.Context,
		arg repository.SelectUserFromProviderNameAndIdParams,
	) (repository.User, error)
	UpdateUserAvatarURL(ctx context.Context, arg repository.UpdateUserAvatarURLParams) error
	CreateUser(ctx context.Context, arg repository.CreateUserParams) (repository.User, error)
}

// Services interface that deals with the todos table
type TasksServices interface {
	SelectAllTasksFromUser(ctx context.Context, userID pgtype.UUID) ([]repository.Todo, error)
	PostTask(ctx context.Context, arg repository.PostTaskParams) error
}
