import DealCard from "../events/DealCard.js";

export default {
  name: "play-card-prince",

  available: () => false,

  perform: function* (state, { targetPlayerId }) {
    const otherPlayer = state.players[targetPlayerId];

    if (!otherPlayer) {
      throw new Error(`Player ${targetPlayerId} not found`);
    }

    yield {
      action: "discard-card",
      playerId: targetPlayerId,
      card: state.players[targetPlayerId].hand[0],
    };

    yield DealCard.create(targetPlayerId, state.deck[0]);
  },
};
