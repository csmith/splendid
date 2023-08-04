import Server from "./server/Server.js";
import express from 'express';
import { createServer } from 'http';
import { handler } from '../build/handler.js'

const port = 3000;
const app = express();
const httpServer = createServer(app);

const server = new Server();
server.bind(httpServer);

app.use(handler);

httpServer.listen(port)