-- name: GetTodoByID :many
    select "id", "user_id", "todo_text", "done", "created_at", "updated_at" from "todos" where "todos"."id" = $1;
-- name: SelectUserBySessionID :one
select "users"."id", "users"."username", "users"."avatar_url", "users"."provider_user_id", "users"."provider_name", "session"."id", "session"."user_id", "session"."expires_at" from "session" inner join "users" on "users"."id" = "session"."user_id" where "session"."id" = $1 limit 1;
-- name: CreateSession :exec
    insert into "session" ("id", "user_id", "expires_at") values ($1, $2, $3);
