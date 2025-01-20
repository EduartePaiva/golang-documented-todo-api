import { existsSync, mkdirSync, writeFileSync } from "fs";
import path from "path";
import inserts from "./queries/inserts";
import selects from "./queries/selects";

export function buildQueries() {
    const dir = path.join(__dirname, "../sqlc");
    if (!existsSync(dir)) {
        mkdirSync(dir, { recursive: true });
    }
    let output = "";

    for (const select of selects) {
        output += select;
    }
    for (const insert of inserts) {
        output += insert;
    }

    writeFileSync(path.join(__dirname, "../sqlc/query.sql"), output);
}
