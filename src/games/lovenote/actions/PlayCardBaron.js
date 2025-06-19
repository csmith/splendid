import CardNoOp from "../events/CardNoOp.js";
import HandRevealed from "../events/HandRevealed.js";
import { areAllProtected } from "../util.js";
import EliminatePlayer from "./EliminatePlayer.js";

export default {
  name: "play-card-baron",

  available: () => false,

  perform: function* (state, { playerData, targetPlayerId }) {
    // If all other players are protected, it does nothing
    if (areAllProtected(state, playerData.details.id)) {
      yield CardNoOp.create(playerData.details.id);
      return;
    }

    const otherPlayer = state.players[targetPlayerId];
    if (!otherPlayer) {
      throw new Error(`Player ${targetPlayerId} not found`);
    }

    yield HandRevealed.create(playerData.details.id, otherPlayer.details.id, otherPlayer.hand);
    yield HandRevealed.create(otherPlayer.details.id, playerData.details.id, playerData.hand);

    if (otherPlayer.hand[0].closeness > playerData.hand[0].closeness) {
      yield* EliminatePlayer.perform(state, {
        playerId: playerData.details.id,
        reason: `${playerData.details.name} deployed a Baron against ${otherPlayer.details.name} and lost the comparison`,
      });
    } else if (otherPlayer.hand[0].closeness < playerData.hand[0].closeness) {
      yield* EliminatePlayer.perform(state, {
        playerId: otherPlayer.details.id,
        reason: `${playerData.details.name} deployed a Baron and won the comparison`,
      });
    }
  },
};
