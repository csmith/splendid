import {Manager} from "socket.io-client";
import {derived, get, writable} from "svelte/store";
import {newPlayer} from "../common/player.js";
import Engine from "../common/engine.js";
import phases from "../splendid/phases.js";

export default class {

    #storage;

    #connected = writable(false);
    #player = writable(null);
    #gameType = writable(null);
    #gameId = writable(null);
    #gameState = writable({phase: ''});

    // TODO: Don't just use the splendid phases :D
    #engine = derived(this.#gameState, $state => new Engine($state, phases));

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
        return derived(this.#engine, $engine => $engine.actions(get(this.#player)));
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

    #onGameAction({name, args, state}) {
        this.#gameState.set(state);
    }

    #onGameJoined({type, id, state}) {
        this.#storage.setItem('game', id);
        this.#gameId.set(id);
        this.#gameState.set(state);
        this.#gameType.set(type);
    }

}