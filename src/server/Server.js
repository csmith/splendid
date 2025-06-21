import { verify } from "../common/crypto.js";
import Engine from "../common/engine.js";
import games from "../games.js";
import SetOptions from "../games/shared/events/SetOptions.js";
import Connection from "./Connection.js";
import { Server } from "socket.io";

export default class {
  /** @type Store */
  #store;

  #io;
  #games = {};

  #sockets = {};

  constructor(store) {
    this.#store = store;

    setInterval(() => this.saveGames(), 1000 * 60);
  }

  bind(server) {
    this.#io = new Server(server, { cors: { origin: "*" } });

    this.#io.of("/game-server").on("connection", (socket) => {
      this.#sockets[socket.id] = { handler: new Connection(this, socket) };
    });
  }

  startGame(name, options = {}) {
    const game = Object.values(games).find((g) => g.name === name);
    if (!game) {
      throw new Error(`Game ${name} not found`);
    }

    const engine = new Engine(game);
    
    if (Object.keys(options).length > 0) {
      engine.applyEvent(SetOptions.create(options));
    }
    
    const id = this.#store.assignId(engine);
    this.#games[id] = engine;
    return { id, engine, game };
  }

  #loadGame(id) {
    if (this.#games[id]) {
      return this.#games[id];
    }

    const engine = this.#store.loadGame(id);
    this.#games[id] = engine;
    return engine;
  }

  saveGames() {
    Object.entries(this.#games).forEach(([id, engine]) => {
      this.#store.saveGame(id, engine);
    });
  }

  joinGame(id, verification, player, socketId) {
    const engine = this.#loadGame(id);
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

    const game = Object.values(games).find((g) => g.name === engine.type);
    return { id, engine, game };
  }
}
