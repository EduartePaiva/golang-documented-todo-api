import { InferInsertModel, InferSelectModel } from "drizzle-orm";
import {
    pgEnum,
    pgTable,
    uniqueIndex,
    uuid,
    varchar,
} from "drizzle-orm/pg-core";

export const providerEnum = pgEnum("provider_name", ["github", "google"]);

const usersTable = pgTable(
    "users",
    {
        id: uuid().defaultRandom().primaryKey().notNull(),
        username: varchar({ length: 255 }).notNull(),
        avatarUrl: varchar({ length: 560 }),
        providerUserId: varchar({ length: 255 }).notNull(),
        providerName: providerEnum().notNull(),
    },
    (table) => [
        uniqueIndex("provider_index").on(
            table.providerName,
            table.providerUserId
        ),
    ]
);

export type User = InferSelectModel<typeof usersTable>;
export type InsertUser = InferInsertModel<typeof usersTable>;

export default usersTable;
