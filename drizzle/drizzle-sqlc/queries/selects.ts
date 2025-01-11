import { eq } from "drizzle-orm";
import db from "../../db";
import { todos } from "../../db/schema";
import { generateSelectQuery } from "./queries-utils";

const selectTodoById = generateSelectQuery(
    "GetTodoByID",
    db.select().from(todos).where(eq(todos.id, "")).toSQL()
);

export default [selectTodoById];
