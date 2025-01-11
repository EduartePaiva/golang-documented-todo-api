import { defineConfig } from "drizzle-kit";

import env from "./drizzle/env";

export default defineConfig({
    schema: "./drizzle/db/schema",
    dialect: "postgresql",
    out: "./drizzle/db/migrations",
    dbCredentials: {
        url: env.DATABASE_URL,
    },
    casing: "snake_case",
});
