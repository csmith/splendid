export default {
  name: "play-card-priest",

  available: () => false,

  perform: function* (state, { playerData, targetPlayerId }) {
    const otherPlayer = state.players[targetPlayerId];

    if (!otherPlayer) {
      throw new Error(`Player ${targetPlayerId} not found`);
    }

    yield {
      event: "hand-revealed",
      hand: otherPlayer.hand,
      handPlayerId: otherPlayer.details.id,
      playerId: playerData.details.id,
    };
  },
};
