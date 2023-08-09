import events from "./events.js";
import masker from "./masker.js";
import phases from "./phases.js";
import state from "./state.js";

export default {
  name: "Love note",
  players: { min: 2, max: 4 },
  based_on: {
    game: "Love Letter",
    creator: "Seiji Kanai",
    link: "https://www.zmangames.com/en/games/love-letter/",
    purchase: "https://www.board-game.co.uk/product/love-letter-z-man-games/",
  },
  description: "Use your connections to beat out the competition and get your note delivered to the Princess.",
  phases,
  state,
  masker,
  events,
};
