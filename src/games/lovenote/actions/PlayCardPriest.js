import HandRevealed from "../events/HandRevealed.js";

export default {
  name: "play-card-priest",

  available: () => false,

  perform: function* (state, { playerData, targetPlayerId }) {
    const otherPlayer = state.players[targetPlayerId];

    if (!otherPlayer) {
      throw new Error(`Player ${targetPlayerId} not found`);
    }

    yield HandRevealed.create(playerData.details.id, otherPlayer.details.id, otherPlayer.hand);
  },
};
