import db from "@/db";
import { sessionTable } from "@/db/schema";
import { eq } from "drizzle-orm";
import { generateExecQuery } from "./queries-utils";

const deleteSessionById = generateExecQuery(
    "DeleteSessionByID",
    db.delete(sessionTable).where(eq(sessionTable.id, "")).toSQL()
);

export default [deleteSessionById];
