import cards from "../../../src/games/lovenote/data/cards.js";
import { Before } from "@cucumber/cucumber";

Before(function () {
  this.parseLoveNoteCard = function (name) {
    return {
      ...cards.find((card) => card.type === name),
      id: crypto.randomUUID(),
    };
  };
});
