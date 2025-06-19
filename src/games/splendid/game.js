import events from "./events.js";
import phases from "./phases.js";
import state from "./state.js";

export default {
  name: "Splendid",
  players: { min: 2, max: 4 },
  based_on: {
    game: "Splendor",
    creator: "Marc André",
    link: "https://www.spacecowboys.fr/splendor-english",
    purchase: "https://www.board-game.co.uk/product/splendor/",
  },
  description:
    "You play as a merchant, buying up developments to gain prestige and attract noble visitors. Do you build up a " +
    "strong foundation to support your future purchases, or do you rush to buy the most prestigious developments " +
    "sooner?",
  phases,
  state,
  events,
};
