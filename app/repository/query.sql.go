// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSession = `-- name: CreateSession :exec
    insert into "session" ("id", "user_id", "expires_at") values ($1, $2, $3)
`

type CreateSessionParams struct {
	ID        string
	UserID    pgtype.UUID
	ExpiresAt pgtype.Timestamptz
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) error {
	_, err := q.db.Exec(ctx, createSession, arg.ID, arg.UserID, arg.ExpiresAt)
	return err
}

const createUser = `-- name: CreateUser :one
    insert into "users" ("id", "username", "avatar_url", "provider_user_id", "provider_name") values (default, $1, $2, $3, $4) returning "id", "username", "avatar_url", "provider_user_id", "provider_name"
`

type CreateUserParams struct {
	Username       string
	AvatarUrl      pgtype.Text
	ProviderUserID string
	ProviderName   ProviderName
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.AvatarUrl,
		arg.ProviderUserID,
		arg.ProviderName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.AvatarUrl,
		&i.ProviderUserID,
		&i.ProviderName,
	)
	return i, err
}

const deleteSessionByID = `-- name: DeleteSessionByID :exec
    delete from "session" where "session"."id" = $1
`

func (q *Queries) DeleteSessionByID(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteSessionByID, id)
	return err
}

const getTodoByID = `-- name: GetTodoByID :many
    select "id", "user_id", "todo_text", "done", "created_at", "updated_at" from "todos" where "todos"."id" = $1
`

func (q *Queries) GetTodoByID(ctx context.Context, id pgtype.UUID) ([]Todo, error) {
	rows, err := q.db.Query(ctx, getTodoByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.TodoText,
			&i.Done,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectUserBySessionID = `-- name: SelectUserBySessionID :one
select "users"."id", "users"."username", "users"."avatar_url", "users"."provider_user_id", "users"."provider_name", "session"."id", "session"."user_id", "session"."expires_at" from "session" inner join "users" on "users"."id" = "session"."user_id" where "session"."id" = $1 limit 1
`

type SelectUserBySessionIDRow struct {
	ID             pgtype.UUID
	Username       string
	AvatarUrl      pgtype.Text
	ProviderUserID string
	ProviderName   ProviderName
	ID_2           string
	UserID         pgtype.UUID
	ExpiresAt      pgtype.Timestamptz
}

func (q *Queries) SelectUserBySessionID(ctx context.Context, id string) (SelectUserBySessionIDRow, error) {
	row := q.db.QueryRow(ctx, selectUserBySessionID, id)
	var i SelectUserBySessionIDRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.AvatarUrl,
		&i.ProviderUserID,
		&i.ProviderName,
		&i.ID_2,
		&i.UserID,
		&i.ExpiresAt,
	)
	return i, err
}

const selectUserFromProviderNameAndId = `-- name: SelectUserFromProviderNameAndId :one
select "id", "username", "avatar_url", "provider_user_id", "provider_name" from "users" where ("users"."provider_user_id" = $1 and "users"."provider_name" = $2) limit 1
`

type SelectUserFromProviderNameAndIdParams struct {
	ProviderUserID string
	ProviderName   ProviderName
}

func (q *Queries) SelectUserFromProviderNameAndId(ctx context.Context, arg SelectUserFromProviderNameAndIdParams) (User, error) {
	row := q.db.QueryRow(ctx, selectUserFromProviderNameAndId, arg.ProviderUserID, arg.ProviderName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.AvatarUrl,
		&i.ProviderUserID,
		&i.ProviderName,
	)
	return i, err
}

const updateSessionExpiresAt = `-- name: UpdateSessionExpiresAt :exec
    update "session" set "expires_at" = $1 where "session"."id" = $2
`

type UpdateSessionExpiresAtParams struct {
	ExpiresAt pgtype.Timestamptz
	ID        string
}

func (q *Queries) UpdateSessionExpiresAt(ctx context.Context, arg UpdateSessionExpiresAtParams) error {
	_, err := q.db.Exec(ctx, updateSessionExpiresAt, arg.ExpiresAt, arg.ID)
	return err
}

const updateUserAvatarURL = `-- name: UpdateUserAvatarURL :exec
    update "users" set "avatar_url" = $1 where "users"."id" = $2
`

type UpdateUserAvatarURLParams struct {
	AvatarUrl pgtype.Text
	ID        pgtype.UUID
}

func (q *Queries) UpdateUserAvatarURL(ctx context.Context, arg UpdateUserAvatarURLParams) error {
	_, err := q.db.Exec(ctx, updateUserAvatarURL, arg.AvatarUrl, arg.ID)
	return err
}
