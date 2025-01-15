import { config } from "dotenv";
import { expand } from "dotenv-expand";
import z from "zod";

expand(config());

const EnvSchema = z.object({
    NODE_ENV: z.enum(["development", "production", "test"]),
    DATABASE_URL: z.string().url(),
    DB_HOST: z.string(),
    DB_USER: z.string(),
    DB_PASSWORD: z.string(),
    DB_NAME: z.string(),
    DB_PORT: z.coerce.number(),
    DB_MIGRATING: z
        .string()
        .refine((s) => s === "true" || s === "false")
        .transform((s) => s === "true")
        .optional(),
});
const parsedData = EnvSchema.safeParse(process.env);
if (!parsedData.success) {
    console.error(
        "❌ Invalid environment variables:",
        parsedData.error.flatten().fieldErrors
    );
    process.exit(1);
}
const env = parsedData.data;
export default env;
