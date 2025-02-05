-- name: GetTodoByID :many
    select "id", "user_id", "todo_text", "done", "created_at", "updated_at" from "todos" where "todos"."id" = $1;
-- name: SelectUserBySessionID :one
select "users"."id", "users"."username", "users"."avatar_url", "users"."provider_user_id", "users"."provider_name", "session"."id", "session"."user_id", "session"."expires_at" from "session" inner join "users" on "users"."id" = "session"."user_id" where "session"."id" = $1 limit 1;
-- name: SelectUserFromProviderNameAndId :one
select "id", "username", "avatar_url", "provider_user_id", "provider_name" from "users" where ("users"."provider_user_id" = $1 and "users"."provider_name" = $2) limit 1;
-- name: SelectAllTasksFromUser :many
    select "id", "user_id", "todo_text", "done", "created_at", "updated_at" from "todos" where "todos"."user_id" = $1 order by "todos"."created_at" desc;
-- name: CreateSession :exec
    insert into "session" ("id", "user_id", "expires_at") values ($1, $2, $3);
-- name: CreateUser :one
    insert into "users" ("id", "username", "avatar_url", "provider_user_id", "provider_name") values (default, $1, $2, $3, $4) returning "id", "username", "avatar_url", "provider_user_id", "provider_name";
-- name: PostTask :exec
    insert into "todos" ("id", "user_id", "todo_text", "done", "created_at", "updated_at") values ($1, $2, $3, $4, default, $5) on conflict ("id","user_id") do update set "todo_text" = $6, "done" = $7, "updated_at" = $8;
-- name: UpdateSessionExpiresAt :exec
    update "session" set "expires_at" = $1 where "session"."id" = $2;
-- name: UpdateUserAvatarURL :exec
    update "users" set "avatar_url" = $1 where "users"."id" = $2;
-- name: UpdateTextFromTodo :exec
    update "todos" set "todo_text" = $1 where ("todos"."id" = $2 and "todos"."user_id" = $3);
-- name: UpdateDoneFromTodo :exec
    update "todos" set "done" = $1 where ("todos"."id" = $2 and "todos"."user_id" = $3);
-- name: UpdateDoneAndTextFromTodo :exec
    update "todos" set "todo_text" = $1, "done" = $2 where ("todos"."id" = $3 and "todos"."user_id" = $4);
-- name: DeleteSessionByID :exec
    delete from "session" where "session"."id" = $1;
-- name: DeleteTaskByIDAndUserID :exec
    delete from "todos" where ("todos"."id" = $1 and "todos"."user_id" = $2);
