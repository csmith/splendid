import DealCard from "../events/DealCard.js";
import DiscardCard from "./DiscardCard.js";

export default {
  name: "play-card-prince",

  available: () => false,

  perform: function* (state, { targetPlayerId }) {
    const otherPlayer = state.players[targetPlayerId];

    if (!otherPlayer) {
      throw new Error(`Player ${targetPlayerId} not found`);
    }

    yield* DiscardCard.perform(state, {
      playerId: targetPlayerId,
      card: state.players[targetPlayerId].hand[0],
    });

    yield DealCard.create(targetPlayerId, state.deck[0]);
  },
};
