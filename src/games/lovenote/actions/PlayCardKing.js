import CardNoOp from "../events/CardNoOp.js";
import SwapHands from "../events/SwapHands.js";
import { areAllProtected } from "../util.js";

export default {
  name: "play-card-king",

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

    yield SwapHands.create([
      {
        id: targetPlayerId,
        hand: playerData.hand,
      },
      {
        id: playerData.details.id,
        hand: otherPlayer.hand,
      },
    ]);
  },
};
