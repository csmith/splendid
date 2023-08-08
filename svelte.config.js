import adapter from "@sveltejs/adapter-node";
import preprocess from "svelte-preprocess";

/** @type {import('@sveltejs/kit').Config} */
const config = {
  preprocess: preprocess(),
  kit: {
    adapter: adapter(),
    files: {
      assets: "static",
      routes: "src/frontend/routes",
      appTemplate: "src/frontend/app.html",
    },
  },
};

export default config;
