export default {
  name: "eliminate-player",

  available: () => false,

  perform: function* (state, { playerId, reason }) {
    yield {
      event: "eliminate-player",
      playerId,
      reason,
    };

    const playerData = state.players[playerId];
    yield* playerData.hand.map((c) => ({
      event: "discard-card",
      playerId,
      card: c,
    }));
  },
};
