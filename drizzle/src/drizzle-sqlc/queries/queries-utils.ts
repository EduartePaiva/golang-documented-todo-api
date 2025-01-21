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

/**
 * This function have to contains a limit on it otherwise it'll throw
 */
export function generateSelectOneQuery(name: string, query: Query): string {
    const queryString = query.sql;
    if (!queryString.includes("limit")) {
        throw new Error("this function should include a limit");
    }
    const result = queryString.replace(
        /limit (\$[0-9]+)/g,
        (value) => value.slice(0, 6) + "1"
    );
    return `-- name: ${name} :one\n${result};\n`;
}
