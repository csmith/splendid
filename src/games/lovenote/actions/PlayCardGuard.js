import CardNoOp from "../events/CardNoOp.js";
import GuardFailed from "../events/GuardFailed.js";
import { areAllProtected } from "../util.js";
import EliminatePlayer from "./EliminatePlayer.js";

export default {
  name: "play-card-guard",

  available: () => false,

  perform: function* (state, { playerData, targetPlayerId, guessedType }) {
    // If all other players are protected, it does nothing
    if (areAllProtected(state, playerData.details.id)) {
      yield CardNoOp.create(playerData.details.id);
      return;
    }

    const otherPlayer = state.players[targetPlayerId];
    if (!otherPlayer) {
      throw new Error(`Player ${targetPlayerId} not found`);
    }

    if (otherPlayer.hand.find((c) => c.type === guessedType)) {
      yield* EliminatePlayer.perform(state, {
        playerId: otherPlayer.details.id,
        reason: `${playerData.details.name} deployed a Guard and correctly guessed they held a ${guessedType}`,
      });
    } else {
      yield GuardFailed.create(playerData.details.id);
    }
  },
};
