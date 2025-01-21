import db from "@/db";
import { sessionTable } from "@/db/schema";
import usersTable from "@/db/schema/users";
import { eq } from "drizzle-orm";
import { describe, expect, it } from "vitest";
import { generateSelectOneQuery } from "./queries-utils";

describe("Test the generateSelectOneQuery", () => {
    it("Test if the limit is properly handled", () => {
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
        const expectResult = `-- name: SelectUserBySessionID :one\nselect "users"."id", "users"."username", "users"."avatar_url", "users"."provider_user_id", "users"."provider_name", "session"."id", "session"."user_id", "session"."expires_at" from "session" inner join "users" on "users"."id" = "session"."user_id" where "session"."id" = $1 limit 1;\n`;
        expect(selectUserBySessionID).toBe(expectResult);
    });

    it("Function should throw if a limit is not specified", () => {
        expect(() =>
            generateSelectOneQuery("Test", db.select().from(usersTable).toSQL())
        ).toThrowError("this function should include a limit");
    });
});
