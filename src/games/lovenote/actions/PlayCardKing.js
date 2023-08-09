import { areAllProtected } from "../util.js";

export default {
  name: "play-card-king",

  available: () => false,

  perform: function* (state, { playerData, targetPlayerId }) {
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

    yield {
      event: "swap-hands",
      players: [
        {
          id: targetPlayerId,
          hand: playerData.hand,
        },
        {
          id: playerData.details.id,
          hand: otherPlayer.hand,
        },
      ],
    };
  },
};
