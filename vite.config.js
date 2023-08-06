import { sveltekit } from '@sveltejs/kit/vite'
import Server from "./src/server/Server.js";

let server;

const websocketPlugin = {
    name: 'webSocketServer',
    configureServer(v) {
        server = new Server('./data/');
        server.bind(v.httpServer);
    },
    closeBundle: () => server?.saveGames(),
    handleHotUpdate: () => server?.saveGames(),
}

/** @type {import('vite').UserConfig} */
const config = {
    plugins: [sveltekit(), websocketPlugin]
}

export default config