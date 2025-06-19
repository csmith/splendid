import cards from "../data/cards.js";
import { mustDiscardCountess } from "../util.js";
import DiscardCard from "./DiscardCard.js";
import EndTurn from "./EndTurn.js";
import PlayCardBaron from "./PlayCardBaron.js";
import PlayCardCountess from "./PlayCardCountess.js";
import PlayCardGuard from "./PlayCardGuard.js";
import PlayCardHandmaid from "./PlayCardHandmaid.js";
import PlayCardKing from "./PlayCardKing.js";
import PlayCardPriest from "./PlayCardPriest.js";
import PlayCardPrince from "./PlayCardPrince.js";
import PlayCardPrincess from "./PlayCardPrincess.js";

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

    yield* DiscardCard.perform(state, { playerId: player.id, card });

    const cardActions = {
      baron: PlayCardBaron,
      countess: PlayCardCountess,
      guard: PlayCardGuard,
      handmaid: PlayCardHandmaid,
      king: PlayCardKing,
      priest: PlayCardPriest,
      prince: PlayCardPrince,
      princess: PlayCardPrincess,
    };

    const cardAction = cardActions[card.type.toLowerCase()];
    if (cardAction) {
      yield* cardAction.perform(state, { playerData, card, targetPlayerId, guessedType });
    }

    yield* EndTurn.perform(state);
  },
};
