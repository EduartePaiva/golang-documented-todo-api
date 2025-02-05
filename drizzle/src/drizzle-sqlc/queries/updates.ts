import db from "@/db";
import { sessionTable, todos, users } from "@/db/schema";
import { and, eq } from "drizzle-orm";
import { generateExecQuery } from "./queries-utils";

const query = db
    .update(sessionTable)
    .set({
        expiresAt: new Date(),
    })
    .where(eq(sessionTable.id, ""))
    .toSQL();

const updateSessionExpiresAt = generateExecQuery(
    "UpdateSessionExpiresAt",
    query
);

const updateUserAvatarURL = generateExecQuery(
    "UpdateUserAvatarURL",
    db.update(users).set({ avatarUrl: "" }).where(eq(users.id, "")).toSQL()
);

const updateTextFromTask = generateExecQuery(
    "UpdateTextFromTask",
    db
        .update(todos)
        .set({ todoText: "" })
        .where(and(eq(todos.id, ""), eq(todos.userId, "")))
        .toSQL()
);
const updateDoneFromTask = generateExecQuery(
    "UpdateDoneFromTask",
    db
        .update(todos)
        .set({ done: false })
        .where(and(eq(todos.id, ""), eq(todos.userId, "")))
        .toSQL()
);
const updateDoneAndTextFromTask = generateExecQuery(
    "UpdateDoneAndTextFromTask",
    db
        .update(todos)
        .set({ done: false, todoText: "" })
        .where(and(eq(todos.id, ""), eq(todos.userId, "")))
        .toSQL()
);

export default [
    updateSessionExpiresAt,
    updateUserAvatarURL,
    updateTextFromTask,
    updateDoneFromTask,
    updateDoneAndTextFromTask,
];
