import { areAllProtected } from "../util.js";

export default {
  name: "play-card-baron",

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
      event: "hand-revealed",
      hand: otherPlayer.hand,
      handPlayerId: otherPlayer.details.id,
      playerId: playerData.details.id,
    };

    yield {
      event: "hand-revealed",
      hand: playerData.hand,
      handPlayerId: playerData.details.id,
      playerId: otherPlayer.details.id,
    };

    if (otherPlayer.hand[0].closeness > playerData.hand[0].closeness) {
      yield {
        action: "eliminate-player",
        args: {
          playerId: playerData.details.id,
          reason: `${playerData.details.name} deployed a Baron against ${otherPlayer.details.name} and lost the comparison`,
        },
      };
    } else if (otherPlayer.hand[0].closeness < playerData.hand[0].closeness) {
      yield {
        action: "eliminate-player",
        args: {
          playerId: otherPlayer.details.id,
          reason: `${playerData.details.name} deployed a Baron and won the comparison`,
        },
      };
    }
  },
};
