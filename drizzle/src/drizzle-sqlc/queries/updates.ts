import db from "@/db";
import { sessionTable, users } from "@/db/schema";
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

const updateUserAvatarURL = generateExecQuery(
    "UpdateUserAvatarURL",
    db.update(users).set({ avatarUrl: "" }).where(eq(users.id, "")).toSQL()
);

export default [updateSessionExpiresAt, updateUserAvatarURL];
