import { handler } from "../build/handler.js";
import Server from "./server/Server.js";
import Store from "./server/Store.js";
import express from "express";
import { createServer } from "http";

const port = 3000;
const app = express();
const httpServer = createServer(app);

const server = new Server(new Store());

function exitHandler(code) {
  server.saveGames();
  process.exit(code);
}

process.on("SIGQUIT", exitHandler);
process.on("SIGINT", exitHandler);
process.on("SIGTERM", exitHandler);

server.bind(httpServer);

app.use(handler);

httpServer.listen(port);
