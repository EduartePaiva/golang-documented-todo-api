import db from "@/db";
import { sessionTable, todos, users } from "@/db/schema";
import { generateExecQuery, generateInsertOneQuery } from "./queries-utils";

const createSession = generateExecQuery(
    "CreateSession",
    db
        .insert(sessionTable)
        .values({ expiresAt: new Date(), id: "", userId: "" })
        .toSQL()
);

const createUser = generateInsertOneQuery(
    "CreateUser",
    db
        .insert(users)
        .values({
            providerName: "github",
            providerUserId: "",
            username: "",
            avatarUrl: "",
        })
        .returning()
        .toSQL()
);

const postTasks = generateExecQuery(
    "PostTask",
    db
        .insert(todos)
        .values({
            todoText: "",
            userId: "",
            done: false,
            id: "",
            updatedAt: new Date(),
        })
        .onConflictDoUpdate({
            target: [todos.id, todos.userId],
            set: {
                todoText: "",
                done: false,
                updatedAt: new Date(),
            },
        })
        .toSQL()
);

export default [createSession, createUser, postTasks];
