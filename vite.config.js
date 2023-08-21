import Server from "./src/server/Server.js";
import Store from "./src/server/Store.js";
import { sveltekit } from "@sveltejs/kit/vite";

let server;

const websocketPlugin = {
  name: "webSocketServer",
  configureServer(v) {
    server = new Server(new Store());
    server.bind(v.httpServer);
  },
  closeBundle: () => server?.saveGames(),
  handleHotUpdate: () => server?.saveGames(),
};

/** @type {import('vite').UserConfig} */
const config = {
  plugins: [sveltekit(), websocketPlugin],
};

export default config;
