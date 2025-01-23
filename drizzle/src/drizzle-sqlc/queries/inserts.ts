import db from "@/db";
import { sessionTable, users } from "@/db/schema";
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

export default [createSession, createUser];
