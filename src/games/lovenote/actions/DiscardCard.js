import DiscardCard from "../events/DiscardCard.js";

export default {
  name: "discard-card",

  available: () => false,

  perform: function* (state, { playerId, card }) {
    yield DiscardCard.create(playerId, card);

    if (card.type === "Princess" && !state.players[playerId].eliminated) {
      yield {
        action: "eliminate-player",
        playerId,
        reason: `Discarded the Princess`,
      };
    }
  },
};
