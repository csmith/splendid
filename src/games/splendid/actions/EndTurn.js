import { isLastPlayer, nextPlayer } from "../../../common/state.js";
import { canReceiveNoble } from "../util.js";
import _ from "lodash";

export default {
  name: "end-turn",

  available: () => false,

  perform: function* (state) {
    const currentPlayer = state.players[state.turn];
    const finalRound = state.finalRound || Object.values(state.players).some((p) => p.points >= 15);

    if (finalRound && !state.finalRound) {
      yield {
        event: "final-round",
      };
    }

    // First, if there are any nobles that can be received, the player must receive one
    if (state.phase === "play" && state.nobles.some((n) => canReceiveNoble(currentPlayer, n))) {
      yield {
        event: "change-phase",
        phase: "noble",
      };
      return;
    }

    // If this is the final round and this is the last player, the game is over
    if (finalRound && isLastPlayer(state, currentPlayer.details)) {
      yield* [
        {
          event: "change-phase",
          phase: "end",
        },
        {
          event: "change-player",
          playerId: undefined,
        },
      ];
      return;
    }

    // If we're not finishing immediately, and the current player is over their token limit, they must discard
    if (_.sum(Object.values(currentPlayer.tokens)) > 10) {
      yield {
        event: "change-phase",
        phase: "discard",
      };
      return;
    }

    // Otherwise, the next player plays
    if (state.phase !== "play") {
      yield {
        event: "change-phase",
        phase: "play",
      };
    }

    yield {
      event: "change-player",
      playerId: nextPlayer(state),
    };
  },
};
