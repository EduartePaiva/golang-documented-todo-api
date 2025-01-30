import { existsSync, mkdirSync, writeFileSync } from "fs";
import path from "path";
import deletes from "./queries/deletes";
import inserts from "./queries/inserts";
import selects from "./queries/selects";
import updates from "./queries/updates";

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
    for (const update of updates) {
        output += update;
    }
    for (const del of deletes) {
        output += del;
    }

    writeFileSync(path.join(__dirname, "../sqlc/query.sql"), output);
}
