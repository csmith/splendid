import RemoveCardFromDeck from "../events/RemoveCardFromDeck.js";
import ReserveCard from "../events/ReserveCard.js";
import TakeTokens from "../events/TakeTokens.js";
import EndTurn from "./EndTurn.js";

export default {
  name: "reserve-card-from-deck",

  available: function (state, { player }) {
    return player.id === state.turn && state.players[player.id].reserved.length < 3;
  },

  perform: function* (state, { level }) {
    if (level < 1 || level > 3) {
      throw new Error("Invalid level");
    }

    if (state.decks[level - 1].length === 0) {
      throw new Error("Deck is empty");
    }

    const card = state.decks[level - 1][0];
    yield RemoveCardFromDeck.create(level, state.turn, "reserve");
    yield ReserveCard.create(state.turn, card);

    if (state.tokens.gold > 0) {
      yield TakeTokens.create(state.turn, { gold: 1 });
    }

    yield* EndTurn.perform(state);
  },
};
