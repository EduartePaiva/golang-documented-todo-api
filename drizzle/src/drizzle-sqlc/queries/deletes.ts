import db from "@/db";
import { sessionTable, todos } from "@/db/schema";
import { and, eq } from "drizzle-orm";
import { generateExecQuery } from "./queries-utils";

const deleteSessionById = generateExecQuery(
    "DeleteSessionByID",
    db.delete(sessionTable).where(eq(sessionTable.id, "")).toSQL()
);

const deleteTaskByIDAndUserID = generateExecQuery(
    "DeleteTaskByIDAndUserID",
    db
        .delete(todos)
        .where(and(eq(todos.id, ""), eq(todos.userId, "")))
        .toSQL()
);

export default [deleteSessionById, deleteTaskByIDAndUserID];
