import EventEmitter from "events";
import _ from "lodash";

export default class {
  #emitter = new EventEmitter();

  #game;
  #state;
  #events = [];

  constructor(game) {
    this.#game = game;
    this.#state = _.cloneDeep(game?.state);
  }

  onAction(handler) {
    this.#emitter.on("action", handler);
  }

  offAction(handler) {
    this.#emitter.off("action", handler);
  }

  onEvent(handler) {
    this.#emitter.on("event", handler);
  }

  offEvent(handler) {
    this.#emitter.off("event", handler);
  }

  get #phase() {
    return this.#game.phases[this.#state.phase];
  }

  actions(player) {
    return this.#phase?.actions?.filter((a) => a.available(this.#state, { player }));
  }

  #action(name) {
    return _.find(this.#phase.actions, (a) => a.name === name);
  }

  perform(name, player, args) {
    const action = this.#action(name);
    if (!action) {
      throw new Error(`Action ${name} not found`);
    }

    if (!action.available(this.#state, { ...args, player })) {
      throw new Error(`Action ${name} not available`);
    }

    this.#perform(name, { ...args, player });
  }

  #perform(name, args) {
    const action = this.#action(name);
    if (!action) {
      throw new Error(`Action ${name} not found`);
    }

    const result = action.perform(this.#state, args);

    for (let r of result) {
      try {
        if (_.has(r, "event")) {
          this.applyEvent(r);
        } else {
          console.log(`Invalid result of action ${name}: ${JSON.stringify(r)}`);
        }
      } catch (e) {
        console.log(`Failed to process result of action ${name}`, e);
      }
    }
  }

  applyEvent({ event, meta, ...args }) {
    const e = this.#game.events.find((e) => e.name === event);
    if (!e) {
      throw new Error(`Event ${event} not found`);
    }
    
    const eventData = {
      ...args,
      event,
      meta: meta || {
        id: crypto.randomUUID(),
        ts: Date.now(),
      },
    };
    
    e.perform(this.#state, _.cloneDeep(args));
    this.#events.push(eventData);
    this.#emitter.emit("event", eventData);
  }

  get currentPlayer() {
    return this.#state.turn;
  }

  get state() {
    return this.#state;
  }

  get type() {
    return this.#game.name;
  }

  get events() {
    return this.#events;
  }

  // For tests
  set state(state) {
    this.#state = state;
  }
}
