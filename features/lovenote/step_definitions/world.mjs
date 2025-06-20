import cards from "../../../src/games/lovenote/data/cards.js";
import { Before } from "@cucumber/cucumber";

Before(function () {
  this.parseLoveNoteCard = function (name) {
    const card = cards.find((card) => card.type === name);
    if (!card) {
      throw new Error(`Unable to find card type for '${name}'`);
    }
    return {
      ...card,
      id: crypto.randomUUID(),
    };
  };
});
