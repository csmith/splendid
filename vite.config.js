import { sveltekit } from '@sveltejs/kit/vite'
import Server from "./src/server/Server.js";

const websocketPlugin = {
    name: 'webSocketServer',
    configureServer(server) {
        new Server().bind(server.httpServer);
    }
}

/** @type {import('vite').UserConfig} */
const config = {
    plugins: [sveltekit(), websocketPlugin]
}

export default config