import { findPlayer } from "../../../common/state.js";
import { subtractObjects } from "../../../common/util.js";
import AddBonus from "../events/AddBonus.js";
import AddPoints from "../events/AddPoints.js";
import DiscardCard from "../events/DiscardCard.js";
import DiscardReserve from "../events/DiscardReserve.js";
import ReturnTokens from "../events/ReturnTokens.js";
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
      yield ReturnTokens.create(state.turn, deductions);
    }

    if (reserveIndex !== -1) {
      yield DiscardReserve.create(state.turn, card);
    }

    if (index !== -1) {
      yield DiscardCard.create(card, state.turn, "buy");
    }

    if (card.points > 0) {
      yield AddPoints.create(state.turn, card.points);
    }

    yield AddBonus.create(state.turn, card.bonus);

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
