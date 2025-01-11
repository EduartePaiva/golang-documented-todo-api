import { InferInsertModel, InferSelectModel } from "drizzle-orm";
import { boolean, pgTable, text, timestamp, uuid } from "drizzle-orm/pg-core";
import usersTable from "./users";

const todoTable = pgTable("todos", {
    id: uuid().notNull().defaultRandom(),
    userId: uuid()
        .notNull()
        .references(() => usersTable.id),
    todoText: text().notNull(),
    done: boolean().default(false),
    createdAt: timestamp().defaultNow(),
    updatedAt: timestamp().defaultNow(),
});

export type User = InferSelectModel<typeof usersTable>;
export type InsertUser = InferInsertModel<typeof usersTable>;

export default todoTable;
