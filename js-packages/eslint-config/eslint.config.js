import eslintConfigPrettier from "eslint-config-prettier";
import checkFile from "eslint-plugin-check-file";

/** @type {import('eslint').Linter.Config[]} */
export default [
    {
        rules: {
            "no-unused-vars": "warn",
            semi: "error",
        },
    },
    {
        files: ["src/**/*"],
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
