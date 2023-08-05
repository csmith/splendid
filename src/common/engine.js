import _ from "lodash";
import EventEmitter from 'events';

export default class {

    #emitter = new EventEmitter();

    #state;
    #phases;
    #events;
    #masker;
    #type;

    constructor(state, phases, events, masker, type) {
        this.#state = state;
        this.#phases = phases;
        this.#events = events;
        this.#masker = masker;
        this.#type = type;
    }

    onAction(handler) {
        this.#emitter.on("action", handler);
    }

    offAction(handler) {
        this.#emitter.off("action", handler);
    }

    get #phase() {
        return this.#phases[this.#state.phase];
    }

    actions(player) {
        return this.#phase?.actions?.filter((a) => a.available(this.#state, {player}));
    }

    #action(name) {
        return _.find(this.#phase.actions, (a) => a.name === name);
    }

    perform(name, player, args) {
        const action = this.#action(name);
        if (!action) {
            throw new Error(`Action ${name} not found`);
        }

        if (!action.available(this.#state, {...args, player})) {
            throw new Error(`Action ${name} not available`);
        }

        this.#perform(name, {...args, player});
    }

    #perform(name, args) {
        const action = this.#action(name);
        if (!action) {
            throw new Error(`Action ${name} not found`);
        }

        const result = action.perform(this.#state, args);

        _.castArray(result).forEach((r) => {
            try {
                if (!_.has(r, "if") || r.if) {
                    if (_.has(r, "action")) {
                        this.#perform(r.action, r.args);
                    } else if (_.has(r, "event")) {
                        this.#handleEvent(r.event, r);
                    } else {
                        this.#state = r;
                        this.#emitter.emit("action", {name, args, state: r});
                    }
                }
            } catch (e) {
                console.log(`Failed to process result of action ${name}`, e);
            }
        });
    }

    #handleEvent(name, args) {
        const event = this.#events.find((e) => e.name === name);
        if (!event) {
            throw new Error(`Event ${name} not found`);
        }
        event.perform(this.#state, args);
        this.#emitter.emit('event', args);
    }

    stateFor(player, state) {
        return this.#masker(state || this.#state, player);
    }

    get currentPlayer() {
        return this.#state.turn;
    }

    get state() {
        return this.#state;
    }

    get type() {
        return this.#type;
    }

}