import {Server} from "socket.io";
import Splendid from "../splendid/game.js";
import Connection from "./Connection.js";
import Engine from "../common/engine.js";

export default class {

    #io;
    #games = {};

    #availableGames = [
        Splendid
    ];

    #sockets = {};

    bind(server) {
        this.#io = new Server(server, {cors: {origin: '*'}});

        this.#io.of('/game-server').on("connection", (socket) => {
            this.#sockets[socket.id] = {handler: new Connection(this, socket)};
        });
    }

    startGame(name) {
        const game = this.#availableGames.find(game => game.name === name);
        if (!game) {
            throw new Error(`Game ${name} not found`);
        }

        const id = crypto.randomUUID();
        const engine = new Engine(game.state, game.phases, game.masker, game.name);

        this.#games[id] = engine;

        return {id, engine};
    }

    joinGame(id) {
        const engine = this.#games[id];
        if (!engine) {
            throw new Error(`Game ${id} not found`);
        }
        return {id, engine};
    }

};