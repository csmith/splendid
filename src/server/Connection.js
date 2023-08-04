import _ from "lodash";
import {isUUID} from "../common/util.js";

export default class {

    #socket;

    /** @type {import('./Server.js').default} */
    #server;

    #player;

    #gameId;

    /** @type {import('../common/engine.js').default} */
    #engine;

    constructor(server, socket) {
        this.#server = server;
        this.#socket = socket;

        console.log(`New client: ${socket.id}`);

        this.#socket.onAny((event, ...args) => {
            console.log(`Client ${socket.id} sent event '${event}' with args: ${JSON.stringify(args)}`);
        });

        this.#socket.emit('clientConnected', {socketID: this.#socket.id});

        this.#on('set-player', this.#onClientSetPlayer);
        this.#on('start-game', this.#onClientStartGame);
        this.#on('join-game', this.#onClientJoinGame);
        this.#on('perform-action', this.#onClientPerformAction);
    }

    #on(event, handler) {
        this.#socket.on(event, (...args) => {
            try {
                handler.apply(this, args);
            } catch (e) {
                console.log('Error handling socket event ', event, e);
                this.#socket.emit('error', e.message);
            }
        })
    }

    #requirePlayer() {
        if (!this.#player) {
            throw new Error('Player not set');
        }
    }

    #requireNoPlayer() {
        if (this.#player) {
            throw new Error('Player already set');
        }
    }

    #requireEngine() {
        if (!this.#engine) {
            throw new Error('Game not set');
        }
    }

    #requireNoEngine() {
        if (this.#engine) {
            throw new Error('Game already set');
        }
    }

    #setGame({id, engine}) {
        this.#gameId = id;
        this.#engine = engine;
        this.#engine.onAction((a) => this.#onEngineAction(a));
        this.#socket.emit('game-joined', {id, type: engine.type, state: engine.stateFor(this.#player.id)})
    }

    #onClientJoinGame(code) {
        this.#requirePlayer();
        this.#requireNoEngine();

        this.#setGame(this.#server.joinGame(code, this.#player));
    }

    #onClientPerformAction(name, args) {
        this.#requireEngine();
        this.#requirePlayer();

        this.#engine.perform(name, this.#player, args);
    }

    #onClientSetPlayer({id, name}) {
        this.#requireNoPlayer();

        if (!name || name.size < 1 || name.size > 20 || !/^[a-zA-Z0-9 \-]+$/.test(name)) {
            throw new Error('Invalid name');
        }

        if (!id || !isUUID(id)) {
            throw new Error('Invalid player id');
        }

        this.#player = {id, name};
    }

    #onClientStartGame(name) {
        this.#requirePlayer();
        this.#requireNoEngine();

        this.#setGame(this.#server.startGame(name));
    }

    #onEngineAction({name, args, state}) {
        this.#socket.emit('game-action', {name, args, state: this.#engine.stateFor(this.#player.id, state)});
    }

};