import { Query } from "drizzle-orm";

export function generateSelectQuery(name: string, query: Query): string {
    return `-- name: ${name} :many
    ${query.sql};\n`;
}
export function generateInsertOneQuery(name: string, query: Query): string {
    return `-- name: ${name} :one
    ${query.sql};\n`;
}
export function generateInsertExecQuery(name: string, query: Query): string {
    return `-- name: ${name} :exec
    ${query.sql};\n`;
}
