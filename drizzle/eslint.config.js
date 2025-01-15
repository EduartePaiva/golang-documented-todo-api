import pluginJs from "@eslint/js";
import globals from "globals";
import tseslint from "typescript-eslint";

import myConfig from "@golang-documented-todo-api/eslint-config";

/** @type {import('eslint').Linter.Config[]} */
export default [
    { files: ["**/*.{js,mjs,cjs,ts}"] },
    { languageOptions: { globals: globals.browser } },
    pluginJs.configs.recommended,
    ...tseslint.configs.recommended,
    ...myConfig,
];
