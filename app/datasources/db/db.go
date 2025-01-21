package db

import (
	"context"

	"github.com/golang-documented-todo-api/app/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Database interface {
	CreateSession(ctx context.Context, arg repository.CreateSessionParams) error
	GetTodoByID(ctx context.Context, id pgtype.UUID) ([]repository.Todo, error)
	SelectUserBySessionID(ctx context.Context, id string) (repository.SelectUserBySessionIDRow, error)
	UpdateSessionExpiresAt(ctx context.Context, arg repository.UpdateSessionExpiresAtParams) error
	WithTx(tx pgx.Tx) *repository.Queries
}

func NewDatabase(conn *pgx.Conn) Database {
	queries := repository.New(conn)
	return queries
}
