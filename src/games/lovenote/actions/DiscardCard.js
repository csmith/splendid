import DiscardCard from "../events/DiscardCard.js";
import EliminatePlayer from "./EliminatePlayer.js";

export default {
  name: "discard-card",

  available: () => false,

  perform: function* (state, { playerId, card }) {
    yield DiscardCard.create(playerId, card);

    if (card.type === "Princess" && !state.players[playerId].eliminated) {
      yield* EliminatePlayer.perform(state, {
        playerId,
        reason: `Discarded the Princess`,
      });
    }
  },
};
