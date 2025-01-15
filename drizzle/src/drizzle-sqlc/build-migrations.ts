import journal from "@/db/migrations/meta/_journal.json";
import { existsSync, mkdirSync, readFileSync, writeFileSync } from "fs";
import path from "path";

export function buildMigrations() {
    let output = "";

    for (const entry of journal.entries) {
        const sqlFile = readFileSync(
            path.join(__dirname, `../db/migrations/${entry.tag}.sql`)
        ).toString();
        output += `-- ${entry.tag}\n\n`;
        output += sqlFile;
        output += "\n\n";
    }
    const dir = path.join(__dirname, "../sqlc");
    if (!existsSync(dir)) {
        mkdirSync(dir, { recursive: true });
    }

    writeFileSync(path.join(__dirname, "../sqlc/schema.sql"), output);
}
