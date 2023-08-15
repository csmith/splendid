import { findPlayer } from "../../../common/state.js";
import { subtractObjects } from "../../../common/util.js";
import { canAffordCard, costForCard } from "../util.js";
import _ from "lodash";

export default {
  name: "buy-card",

  available: function (state, { player }) {
    return player.id === state.turn;
  },

  perform: function* (state, { player, card }) {
    const index = _.findIndex(state.cards[card.level - 1], (c) => c.id === card.id);
    const reserveIndex = _.findIndex(state.players[player.id].reserved, (c) => c.id === card.id);

    if (index === -1 && reserveIndex === -1) {
      throw new Error("Card not found");
    }

    const playerData = findPlayer(state, player);
    if (!canAffordCard(playerData, card)) {
      throw new Error("Cannot afford card");
    }

    const cost = costForCard(playerData, card);
    const remaining = subtractObjects(playerData.tokens, cost);
    const missing = _.mapValues(Object.fromEntries(Object.entries(remaining).filter((a) => a[1] < 0)), (a) => -a);
    const missingCount = _.sum(Object.values(missing));

    const deductions = { ...subtractObjects(cost, missing), gold: missingCount };

    if (_.sum(Object.values(deductions)) > 0) {
      yield {
        event: "return-tokens",
        playerId: state.turn,
        tokens: deductions,
      };
    }

    if (reserveIndex !== -1) {
      yield {
        event: "discard-reserve",
        playerId: state.turn,
        card,
      };
    }

    if (index !== -1) {
      yield {
        event: "discard-card",
        reason: "buy",
        playerId: state.turn,
        card,
      };
    }

    if (card.points > 0) {
      yield {
        event: "add-points",
        playerId: state.turn,
        points: card.points,
      };
    }

    yield {
      event: "add-bonus",
      playerId: state.turn,
      type: card.bonus,
    };

    if (index !== -1) {
      yield {
        action: "deal",
        level: card.level,
      };
    }

    yield {
      action: "end-turn",
    };
  },
};
