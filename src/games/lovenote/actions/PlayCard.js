import cards from "../data/cards.js";
import { mustDiscardCountess } from "../util.js";

export default {
  name: "play-card",

  available: function (state, { player }) {
    return player.id === state.turn;
  },

  perform: function* (state, { player, cardId, targetPlayerId, guessedType }) {
    const playerData = state.players[player.id];
    const card = playerData.hand.find((card) => card.id === cardId);
    if (!card) {
      throw new Error(`Card ${cardId} not found in player's hand`);
    }

    if (guessedType && cards.every((c) => c.type !== guessedType)) {
      throw new Error(`Invalid guessed type ${guessedType}`);
    }

    const otherPlayer = state.players[targetPlayerId];

    if (otherPlayer && otherPlayer.protected) {
      throw new Error(`Player ${otherPlayer.details.name} cannot be targeted because they have handmaiden cover`);
    }

    if (mustDiscardCountess(state, player.id) && card.type !== "Countess") {
      throw new Error(`You must discard the Countess`);
    }

    yield {
      action: "discard-card",
      playerId: player.id,
      card,
    };

    yield {
      action: `play-card-${card.type.toLowerCase()}`,
      playerData,
      card,
      targetPlayerId,
      guessedType,
    };

    yield {
      action: `end-turn`,
    };
  },
};
