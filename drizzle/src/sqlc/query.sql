-- name: GetTodoByID :many
    select "id", "user_id", "todo_text", "done", "created_at", "updated_at" from "todos" where "todos"."id" = $1;
-- name: CreateSession :exec
    insert into "session" ("id", "user_id", "expires_at") values ($1, $2, $3);
