import DiscardCard from "../events/DiscardCard.js";
import ReserveCard from "../events/ReserveCard.js";
import TakeTokens from "../events/TakeTokens.js";
import _ from "lodash";

export default {
  name: "reserve-card",

  available: function (state, { player }) {
    return player.id === state.turn && state.players[player.id].reserved.length < 3;
  },

  perform: function* (state, { card }) {
    const index = _.findIndex(state.cards[card.level - 1], (c) => c.id === card.id);
    if (index === -1) {
      throw new Error("Card not found");
    }

    yield DiscardCard.create(card, state.turn, "reserve");
    yield ReserveCard.create(state.turn, card);

    if (state.tokens.gold > 0) {
      yield TakeTokens.create(state.turn, { gold: 1 });
    }

    yield* [
      {
        action: "deal",
        level: card.level,
      },
      {
        action: "end-turn",
      },
    ];
  },
};
