export default {
  name: "discard-card",

  available: () => false,

  perform: function* (state, { playerId, card }) {
    yield {
      event: "discard-card",
      playerId,
      card,
    };

    if (card.type === "Princess" && !state.players[playerId].eliminated) {
      yield {
        action: "eliminate-player",
        args: {
          playerId,
          reason: `Discarded the Princess`,
        },
      };
    }
  },
};
