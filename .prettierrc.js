export default {
  plugins: ["@trivago/prettier-plugin-sort-imports", "prettier-plugin-svelte"],
  importOrderSortSpecifiers: true,
  svelteSortOrder: "options-scripts-styles-markup",
};
