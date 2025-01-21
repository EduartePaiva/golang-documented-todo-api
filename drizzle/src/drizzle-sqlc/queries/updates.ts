import db from "@/db";
import { sessionTable } from "@/db/schema";
import { eq } from "drizzle-orm";
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

export default [updateSessionExpiresAt];
