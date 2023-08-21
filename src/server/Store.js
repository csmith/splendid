import Engine from "../common/engine.js";
import { generateId } from "../common/util.js";
import games from "../games.js";
import fs from "fs";
import path from "path";

export default class Store {
  #dir;

  constructor() {
    this.#dir = process.env.SPLENDID_STORAGE_DIR || "./data/";
  }

  loadGame(id) {
    if (!id.match(/^[a-z0-9-]+$/)) {
      throw new Error(`Invalid game ID`);
    }

    if (fs.existsSync(path.join(this.#dir, `${id}.json`))) {
      const { version, game, events } = JSON.parse(fs.readFileSync(path.join(this.#dir, `${id}.json`)));
      if (version !== 1) {
        throw new Error(`Invalid saved state version ${version}`);
      }

      const engine = new Engine(Object.values(games).find((g) => g.name === game));
      events.forEach((event) => engine.applyEvent(event));
      return engine;
    }

    throw new Error(`Game ${id} not found`);
  }

  assignId(engine) {
    let id = generateId();
    while (fs.existsSync(path.join(this.#dir, `${id}.json`))) {
      console.log(`Game ID ${id} already exists, generating new ID...`);
      id = generateId();
    }
    this.saveGame(id, engine);
    return id;
  }

  saveGame(id, engine) {
    fs.writeFileSync(
      path.join(this.#dir, `${id}.json`),
      JSON.stringify({
        version: 1,
        game: engine.type,
        events: engine.events,
      }),
    );
  }
}
