import { Query } from "drizzle-orm";

export function generateSelectQuery(name: string, query: Query): string {
    return `-- name: ${name} :many
    ${query.sql};\n`;
}
