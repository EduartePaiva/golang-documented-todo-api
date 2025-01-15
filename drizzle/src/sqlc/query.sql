-- name: GetTodoByID :many
    select "id", "user_id", "todo_text", "done", "created_at", "updated_at" from "todos" where "todos"."id" = $1;
