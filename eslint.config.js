import pluginJs from "@eslint/js";
import eslintConfigPrettier from "eslint-config-prettier";
import checkFile from "eslint-plugin-check-file";
import globals from "globals";
import tseslint from "typescript-eslint";

/** @type {import('eslint').Linter.Config[]} */
export default [
    { files: ["**/*.{js,mjs,cjs,ts}"] },
    { languageOptions: { globals: globals.browser } },
    pluginJs.configs.recommended,
    ...tseslint.configs.recommended,
    {
        rules: {
            "no-unused-vars": "warn",
            semi: "error",
        },
    },
    {
        files: ["drizzle/**/*"],
        plugins: {
            "check-file": checkFile,
        },
        rules: {
            "check-file/filename-naming-convention": [
                "error",
                {
                    "**/*.ts": "KEBAB_CASE",
                },
                { ignoreMiddleExtensions: true },
            ],
            "check-file/folder-naming-convention": [
                "error",
                {
                    "**": "KEBAB_CASE",
                },
            ],
        },
    },
    eslintConfigPrettier,
];
