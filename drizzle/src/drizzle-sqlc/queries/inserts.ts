import db from "@/db";
import { sessionTable } from "@/db/schema";
import { generateInsertExecQuery } from "./queries-utils";

const createSession = generateInsertExecQuery(
    "CreateSession",
    db
        .insert(sessionTable)
        .values({ expiresAt: new Date(), id: "", userId: "" })
        .toSQL()
);

export default [createSession];
