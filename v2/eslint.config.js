module.exports = [
  {
    ignores: [".yarn", "artifacts", "cache", "dist", "node_modules", "typechain-types", "docs", "crytic-export", "lib"],
    languageOptions: {
      ecmaVersion: 2020,
      sourceType: "module",
      parser: require("@typescript-eslint/parser"),
    },
    plugins: {
      "@typescript-eslint": require("@typescript-eslint/eslint-plugin"),
    },
  },
];
