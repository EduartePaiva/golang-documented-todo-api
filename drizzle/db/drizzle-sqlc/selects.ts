import { eq, Query } from "drizzle-orm";
import db from "..";
import { todos } from "../schema";

function generateSelectQuery(name: string, query: Query): string {
    return `-- name: ${name} :many
    ${query.sql}`;
}

export const selectTodoById = generateSelectQuery(
    "GetTodoByID",
    db.select().from(todos).where(eq(todos.id, "")).toSQL()
);

console.log(selectTodoById);
