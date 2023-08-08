import { isUUID } from "../common/util.js";

export default class {
  #socket;

  /** @type {import('./Server.js').default} */
  #server;

  #player;

  #gameId;

  /** @type {import('../common/engine.js').default} */
  #engine;

  #masker;

  constructor(server, socket) {
    this.#server = server;
    this.#socket = socket;

    console.log(`New client: ${socket.id}`);

    this.#socket.onAny((event, ...args) => {
      console.log(`Client ${socket.id} sent event '${event}' with args: ${JSON.stringify(args)}`);
    });

    this.#socket.emit("clientConnected", { socketID: this.#socket.id });

    this.#on("set-player", this.#onClientSetPlayer);
    this.#on("start-game", this.#onClientStartGame);
    this.#on("join-game", this.#onClientJoinGame);
    this.#on("perform-action", this.#onClientPerformAction);
  }

  #on(event, handler) {
    this.#socket.on(event, (...args) => {
      try {
        handler.apply(this, args);
      } catch (e) {
        console.log("Error handling socket event ", event, e);
        this.#socket.emit("error", e.message);
      }
    });
  }

  #requirePlayer() {
    if (!this.#player) {
      throw new Error("Player not set");
    }
  }

  #requireNoPlayer() {
    if (this.#player) {
      throw new Error("Player already set");
    }
  }

  #requireEngine() {
    if (!this.#engine) {
      throw new Error("Game not set");
    }
  }

  #requireNoEngine() {
    if (this.#engine) {
      throw new Error("Game already set");
    }
  }

  #setGame({ id, engine, masker }) {
    this.#gameId = id;
    this.#engine = engine;
    this.#masker = masker;
    this.#engine.onAction((a) => this.#onEngineAction(a));
    this.#engine.onEvent((a) => this.#onEngineEvent(a));
    this.#socket.emit("game-joined", {
      id,
      type: engine.type,
      events: engine.events.map((e) => masker(e, this.#player.id)),
    });
  }

  #onClientJoinGame({ code, verification }) {
    this.#requirePlayer();
    this.#requireNoEngine();

    this.#setGame(this.#server.joinGame(code, verification, this.#player, this.#socket.id));
  }

  #onClientPerformAction(name, args) {
    this.#requireEngine();
    this.#requirePlayer();

    this.#engine.perform(name, this.#player, args);
  }

  #onClientSetPlayer({ id, name, publicKey }) {
    this.#requireNoPlayer();

    if (!name || name.size < 1 || name.size > 20 || !/^[a-zA-Z0-9 \-]+$/.test(name)) {
      throw new Error("Invalid name");
    }

    if (!id || !isUUID(id)) {
      throw new Error("Invalid player id");
    }

    if (!publicKey) {
      throw new Error("Invalid public key");
    }

    this.#player = { id, name, publicKey };
  }

  #onClientStartGame(name) {
    this.#requirePlayer();
    this.#requireNoEngine();

    this.#setGame(this.#server.startGame(name));
  }

  #onEngineAction(args) {
    this.#socket.emit("game-action", args);
  }

  #onEngineEvent(args) {
    this.#socket.emit("game-event", this.#masker(args, this.#player.id));
  }
}
