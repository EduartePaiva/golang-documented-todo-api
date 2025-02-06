package db

import (
	"context"

	"github.com/golang-documented-todo-api/app/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	GetTodoByID(ctx context.Context, id pgtype.UUID) ([]repository.Todo, error)
	CreateSession(ctx context.Context, arg repository.CreateSessionParams) error
	SelectUserBySessionID(ctx context.Context, id string) (repository.SelectUserBySessionIDRow, error)
	UpdateSessionExpiresAt(ctx context.Context, arg repository.UpdateSessionExpiresAtParams) error
	SelectUserFromProviderNameAndId(
		ctx context.Context,
		arg repository.SelectUserFromProviderNameAndIdParams,
	) (repository.User, error)
	UpdateUserAvatarURL(ctx context.Context, arg repository.UpdateUserAvatarURLParams) error
	CreateUser(ctx context.Context, arg repository.CreateUserParams) (repository.User, error)
	WithTx(tx pgx.Tx) *repository.Queries
	DeleteSessionByID(ctx context.Context, id string) error
	SelectAllTasksFromUser(ctx context.Context, userID pgtype.UUID) ([]repository.Todo, error)
	PostTask(ctx context.Context, arg repository.PostTaskParams) error
	DeleteTaskByIDAndUserID(ctx context.Context, arg repository.DeleteTaskByIDAndUserIDParams) error
	UpdateDoneAndTextFromTask(ctx context.Context, arg repository.UpdateDoneAndTextFromTaskParams) error
	UpdateDoneFromTask(ctx context.Context, arg repository.UpdateDoneFromTaskParams) error
	UpdateTextFromTask(ctx context.Context, arg repository.UpdateTextFromTaskParams) error
}

func NewDatabase(conn *pgxpool.Pool) Database {
	queries := repository.New(conn)
	return queries
}
