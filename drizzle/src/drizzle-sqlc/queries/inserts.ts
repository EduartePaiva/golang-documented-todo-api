import db from "@/db";
import { sessionTable } from "@/db/schema";
import { generateExecQuery } from "./queries-utils";

const createSession = generateExecQuery(
    "CreateSession",
    db
        .insert(sessionTable)
        .values({ expiresAt: new Date(), id: "", userId: "" })
        .toSQL()
);

export default [createSession];
