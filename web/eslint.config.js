import pluginVue from "eslint-plugin-vue";
import globals from "globals";
import { defineConfig } from "eslint/config";
import js from "@eslint/js";
import css from "@eslint/css";
import eslintConfigPrettier from "eslint-config-prettier/flat";

export default defineConfig([
  // add more generic rulesets here, such as:
  // js.configs.recommended,
  ...pluginVue.configs["flat/recommended"],
  js.configs.recommended,
  css.configs.recommended,
  eslintConfigPrettier,
  {
    rules: {
      // override/add rules settings here, such as:
      // 'vue/no-unused-vars': 'error'
    },
    languageOptions: {
      sourceType: "module",
      globals: {
        ...globals.browser,
      },
    },
  },
]);
