import { buildQueries } from "./drizzle-sqlc/build-queries";

// it seems that sqlc can build the schema directly from the migrations folder
// buildMigrations();
buildQueries();
