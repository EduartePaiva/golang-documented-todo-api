// Make sure to install the 'postgres' package
import { drizzle } from "drizzle-orm/postgres-js";
import postgres from "postgres";

import env from "@/env";

export const client = postgres(env.DATABASE_URL, {
    max: env.DB_MIGRATING ? 1 : undefined,
});

const db = drizzle({ client: client, casing: "snake_case" });
export default db;
