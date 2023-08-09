import { areAllProtected } from "../util.js";

export default {
  name: "play-card-guard",

  available: () => false,

  perform: function* (state, { playerData, targetPlayerId, guessedType }) {
    // If all other players are protected, it does nothing
    if (areAllProtected(state, playerData.details.id)) {
      yield {
        event: "card-no-op",
        playerId: playerData.details.id,
      };
      return;
    }

    const otherPlayer = state.players[targetPlayerId];
    if (!otherPlayer) {
      throw new Error(`Player ${targetPlayerId} not found`);
    }

    if (otherPlayer.hand.find((c) => c.type === guessedType)) {
      yield {
        action: "eliminate-player",
        args: {
          playerId: otherPlayer.details.id,
          reason: `${playerData.details.name} deployed a Guard and correctly guessed they held a ${guessedType}`,
        },
      };
    } else {
      yield {
        event: "guard-failed",
        playerId: playerData.details.id,
      };
    }
  },
};
