import { findPlayer } from "../../common/state.js";
import { addObjects, subtractObjects } from "../../common/util.js";
import _ from "lodash";

export default {
  name: "reserve-card",

  available: function (state, { player }) {
    return player.id === state.turn && state.players[player.id].reserved.length < 3;
  },

  perform: function (state, { card }) {
    const index = _.findIndex(state.cards[card.level - 1], (c) => c.id === card.id);
    if (index === -1) {
      throw new Error("Card not found");
    }

    const getsGold = state.tokens.gold > 0;

    return _.concat(
      {
        event: "discard-card",
        reason: "reserve",
        playerId: state.turn,
        card,
      },
      {
        event: "reserve-card",
        playerId: state.turn,
        card,
      },
      {
        if: getsGold,
        event: "take-tokens",
        playerId: state.turn,
        tokens: { gold: 1 },
      },
      {
        action: "deal",
        args: {
          level: card.level,
        },
      },
      {
        action: "end-turn",
      },
    );
  },
};
