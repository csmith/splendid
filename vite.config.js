import { sveltekit } from '@sveltejs/kit/vite'
import Server from "./src/server/Server.js";

let server = new Server('./data/');

const websocketPlugin = {
    name: 'webSocketServer',
    configureServer(v) {
        server.bind(v.httpServer);
    },
    closeBundle: () => server.saveGames(),
    handleHotUpdate: () => server.saveGames(),
}

/** @type {import('vite').UserConfig} */
const config = {
    plugins: [sveltekit(), websocketPlugin]
}

export default config