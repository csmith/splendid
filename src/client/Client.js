import { createAttestation, getPublicKey, newPrivateKey } from "../common/crypto.js";
import Engine from "../common/engine.js";
import { newPlayer } from "../common/player.js";
import games from "../games.js";
import { Manager } from "socket.io-client";
import { derived, get, writable } from "svelte/store";

export default class {
  #storage;

  #engine;

  #connected = writable(false);
  #player = writable(null);
  #gameType = writable(null);
  #gameId = writable(null);
  #gameState = writable({});
  #gameEvents = writable([]);

  #pendingEvents = [];
  #nextEvent = writable(null);

  #gameToJoin;

  #manager;
  #socket;

  constructor(storage = window.localStorage) {
    this.#storage = storage;
    this.#manager = new Manager(location.protocol === "https:" ? "wss:///" : "ws:///", { autoConnect: false });
    this.#socket = this.#manager.socket("/game-server");

    const savedPlayer = this.#storage.getItem("player");
    if (savedPlayer) {
      this.#player.set(JSON.parse(savedPlayer));
    }

    this.#socket.on("connect", () => this.#onConnect());
    this.#socket.on("disconnect", () => this.#onDisconnect());
    this.#socket.on("game-joined", (args) => this.#onGameJoined(args));
    this.#socket.on("game-action", (args) => this.#onGameAction(args));
    this.#socket.on("game-event", (args) => this.#onGameEvent(args));
    this.#socket.onAny((event, ...args) => {
      console.log(event, args);
    });
  }

  on(event, callback) {
    this.#socket.on(event, callback);
  }

  connect() {
    this.#socket.connect();
  }

  get isInGame() {
    return derived(this.#gameType, ($game) => $game !== null);
  }

  get isConnected() {
    return this.#connected;
  }

  get hasPlayer() {
    return derived(this.#player, ($player) => $player !== null);
  }

  get playerId() {
    return derived(this.#player, ($player) => $player?.id);
  }

  get actions() {
    return derived(this.#gameState, () => {
      const player = get(this.#player);
      if (player) {
        return this.#engine?.actions(player)?.map((a) => a.name);
      } else {
        return [];
      }
    });
  }

  get gameId() {
    return this.#gameId;
  }

  get gameState() {
    return this.#gameState;
  }

  get gameEvents() {
    return this.#gameEvents;
  }

  get gameType() {
    return this.#gameType;
  }

  createPlayer(displayName) {
    const key = newPrivateKey();
    this.#storage.setItem("player-key", key);
    this.#setPlayer(newPlayer(displayName, getPublicKey(key)));
  }

  #setPlayer(player) {
    this.#player.set(player);
    this.#storage.setItem("player", JSON.stringify(player));
    this.#socket.emit("set-player", player);
    // Bop the store as things may change now we know who we are
    this.#gameState.set(get(this.#gameState));

    if (this.#gameToJoin) {
      this.joinGame(this.#gameToJoin);
    }
  }

  startGame(game, options = {}) {
    this.#socket.emit("start-game", { game, options });
  }

  joinGame(code) {
    // When joining an existing game, there may already be a player with the same ID.
    // To prevent other people from impersonating that player, we sign a small chunk of data with our private key.
    // The server can then verify that we have the private key for the player ID we claim to have.
    // If the server hasn't seen the player before, it will store the given public key for future verification.
    this.#socket.emit("join-game", {
      code,
      verification: createAttestation(
        { code, ts: Date.now(), id: this.#socket.id },
        this.#storage.getItem("player-key"),
      ),
    });
  }

  perform(action, args) {
    this.#socket.emit("perform-action", action, { ...args, player: get(this.#player) });
  }

  #onConnect() {
    this.#connected.set(true);

    const player = get(this.#player);
    if (player) {
      this.#setPlayer(player);
    }
  }

  #onDisconnect() {
    this.#connected.set(false);
  }

  #onGameAction({ state }) {
    this.#gameState.set(state);
  }

  #onGameEvent(args) {
    this.#pendingEvents.push(args);
    if (this.#pendingEvents.length === 1) {
      this.#nextEvent.set(args);
    }
  }

  #onGameJoined({ type, id, events }) {
    this.#engine = new Engine(games[type]);
    this.#gameState.set(this.#engine.state);
    this.#gameId.set(id);
    this.#gameType.set(type);
    events.forEach((e) => this.#processEvent(e));
  }

  advanceEvents() {
    const event = this.#pendingEvents.shift();
    if (event) {
      this.#processEvent(event);
      this.#nextEvent.set(this.#pendingEvents[0]);
    }
  }

  #processEvent(e) {
    this.#engine.applyEvent(e);
    this.#gameEvents.set(get(this.#gameEvents).concat([e]));
    this.#gameState.set(this.#engine.state);
  }

  get nextEvent() {
    return this.#nextEvent;
  }

  set game(value) {
    this.#gameToJoin = value;
    if (this.#socket.connected && get(this.#player)) {
      this.joinGame(value);
    }
  }
}
