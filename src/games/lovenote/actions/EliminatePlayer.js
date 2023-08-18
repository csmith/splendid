import DiscardCard from "../events/DiscardCard.js";
import EliminatePlayer from "../events/EliminatePlayer.js";

export default {
  name: "eliminate-player",

  available: () => false,

  perform: function* (state, { playerId, reason }) {
    yield EliminatePlayer.create(playerId, reason);

    const playerData = state.players[playerId];
    yield* playerData.hand.map((c) => DiscardCard.create(playerId, c));
  },
};
