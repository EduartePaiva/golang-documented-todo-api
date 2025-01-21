import db from "@/db";
import { sessionTable, todos } from "@/db/schema";
import usersTable from "@/db/schema/users";
import { eq } from "drizzle-orm";
import { generateSelectOneQuery, generateSelectQuery } from "./queries-utils";

const selectTodoById = generateSelectQuery(
    "GetTodoByID",
    db.select().from(todos).where(eq(todos.id, "")).toSQL()
);

const selectUserBySessionID = generateSelectOneQuery(
    "SelectUserBySessionID",
    db
        .select({ user: usersTable, session: sessionTable })
        .from(sessionTable)
        .where(eq(sessionTable.id, ""))
        .innerJoin(usersTable, eq(usersTable.id, sessionTable.userId))
        .limit(1)
        .toSQL()
);

export default [selectTodoById, selectUserBySessionID];
