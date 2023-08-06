import {Server} from "socket.io";
import Splendid from "../splendid/game.js";
import Connection from "./Connection.js";
import Engine from "../common/engine.js";
import {verify} from "../common/crypto.js";
import fs from "fs";
import path from "path";

export default class {

    #storageDir;

    #io;
    #games = {};

    #availableGames = [
        Splendid
    ];

    #sockets = {};

    constructor(storageDir) {
        this.#storageDir = storageDir || './data';

        setInterval(() => this.saveGames(), 1000 * 60);
    }

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
        const engine = new Engine(game);

        this.#games[id] = engine;

        return {id, engine, masker: game.masker};
    }

    #loadGame(id) {
        if (this.#games[id]) {
            return this.#games[id];
        }

        if (fs.existsSync(path.join(this.#storageDir, `${id}.json`))) {
            const {version, game, events} = JSON.parse(fs.readFileSync(path.join(this.#storageDir, `${id}.json`)));
            if (version !== 1) {
                throw new Error(`Invalid saved state version ${version}`);
            }

            const engine = new Engine(this.#availableGames.find(g => g.name === game));
            events.forEach(event => engine.applyEvent(event));

            this.#games[id] = engine;
            return engine;
        }

        throw new Error(`Game ${id} not found`);
    }

    saveGames() {
        Object.entries(this.#games).forEach(([id, engine]) => {
            fs.writeFileSync(path.join(this.#storageDir, `${id}.json`), JSON.stringify({
                version: 1,
                game: engine.type,
                events: engine.events
            }));
        });
    }

    joinGame(id, verification, player, socketId) {
        const engine = this.#loadGame(id);
        if (!engine) {
            throw new Error(`Game ${id} not found`);
        }

        if (engine.state.players[player.id]) {
            // Make sure the data that's being signed is valid
            if (id !== verification.payload.code) {
                throw new Error(`Invalid game ID in verification payload`);
            }
            if (socketId !== verification.payload.id) {
                throw new Error(`Invalid socket ID in verification payload`);
            }
            if (Math.abs(verification.payload.timestamp - Date.now()) > 1000 * 60 * 5) {
                throw new Error(`Invalid timestamp in verification payload`);
            }

            // And make sure it's signed with the right key
            const publicKey = engine.state.players[player.id].details.publicKey;
            if (!verify(verification.payload, verification.signature, publicKey)) {
                throw new Error(`Invalid signature in verification payload`);
            }
        }

        const game = this.#availableGames.find(game => game.name === engine.type);
        return {id, engine, masker: game.masker};
    }

};