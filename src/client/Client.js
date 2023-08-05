import {Manager} from "socket.io-client";
import {derived, get, writable} from "svelte/store";
import {newPlayer} from "../common/player.js";
import Engine from "../common/engine.js";
import game from "../splendid/game.js";

export default class {

    #storage;

    // TODO: Don't just use the splendid game :D
    #engine = new Engine(game);

    #connected = writable(false);
    #player = writable(null);
    #gameType = writable(null);
    #gameId = writable(null);
    #gameState = writable(this.#engine.state);

    #manager;
    #socket;

    constructor(storage = window.localStorage) {
        this.#storage = storage;
        this.#manager = new Manager('ws:///', {autoConnect: false});
        this.#socket = this.#manager.socket('/game-server');

        this.#socket.on('connect', () => this.#onConnect());
        this.#socket.on('disconnect', () => this.#onDisconnect());
        this.#socket.on('game-joined', (args) => this.#onGameJoined(args));
        this.#socket.on('game-action', (args) => this.#onGameAction(args));
        this.#socket.on('game-event', (args) => this.#onGameEvent(args));
    }

    on(event, callback) {
        this.#socket.on(event, callback);
    }

    connect() {
        this.#socket.connect();

        this.#socket.onAny((event, ...args) => {
            console.log(event, args);
        });
    }

    get isInGame() {
        return derived(this.#gameId, $game => $game !== null);
    }

    get isConnected() {
        return this.#connected;
    }

    get hasPlayer() {
        return derived(this.#player, $player => $player !== null);
    }

    get playerId() {
        return derived(this.#player, $player => $player?.id);
    }

    get actions() {
        return derived(this.#gameState, () => {
            const player = get(this.#player);
            if (player) {
                return this.#engine.actions(player).map((a) => a.name);
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

    get gameType() {
        return this.#gameType;
    }

    createPlayer(displayName) {
        this.#setPlayer(newPlayer(displayName));
    }

    #setPlayer(player) {
        this.#player.set(player);
        this.#storage.setItem('player', JSON.stringify(player));
        this.#socket.emit('set-player', player);
        // Bop the store as things may change now we know who we are
        this.#gameState.set(get(this.#gameState));
    }

    startGame(game) {
        this.#socket.emit('start-game', game);
    }

    joinGame(code) {
        this.#socket.emit('join-game', code);
    }

    perform(action, args) {
         this.#socket.emit('perform-action', action, {...args, player: get(this.#player)});
    }

    #onConnect() {
        this.#connected.set(true);
        if (this.#storage.getItem('player')) {
            this.#setPlayer(JSON.parse(this.#storage.getItem('player')));
        }
        if (this.#storage.getItem('game')) {
            this.#socket.emit('join-game', this.#storage.getItem('game'));
        }
    }

    #onDisconnect() {
        this.#connected.set(false);
    }

    #onGameAction({state}) {
        this.#gameState.set(state);
    }

    #onGameEvent(args) {
        console.log('onGameEvent', args);
        this.#engine.applyEvent(args);
        this.#gameState.set(this.#engine.state);
    }

    #onGameJoined({type, id, events}) {
        this.#storage.setItem('game', id);
        this.#gameId.set(id);
        this.#gameType.set(type);
        events.forEach((e) => this.#onGameEvent(e));
    }

}