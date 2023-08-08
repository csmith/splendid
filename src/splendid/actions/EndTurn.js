import { isLastPlayer, nextPlayer } from "../../common/state.js";
import { canReceiveNoble } from "../util.js";
import _ from "lodash";

export default {
  name: "end-turn",

  available: () => false,

  perform: (state) => {
    const currentPlayer = state.players[state.turn];
    const finalRound = state.finalRound || Object.values(state.players).some((p) => p.points >= 15);

    // First, if there are any nobles that can be received, the player must receive one
    if (state.phase === "play" && state.nobles.some((n) => canReceiveNoble(currentPlayer, n))) {
      return [
        {
          if: finalRound && !state.finalRound,
          event: "final-round",
        },
        {
          event: "change-phase",
          phase: "noble",
        },
      ];
    }

    // If this is the final round and this is the last player, the game is over
    if (finalRound && isLastPlayer(state, currentPlayer.details)) {
      return [
        {
          event: "change-phase",
          phase: "end",
        },
        {
          event: "change-player",
          playerId: undefined,
        },
      ];
    }

    // If we're not finishing immediately, and the current player is over their token limit, they must discard
    if (_.sum(Object.values(currentPlayer.tokens)) > 10) {
      return [
        {
          if: finalRound && !state.finalRound,
          event: "final-round",
        },
        {
          event: "change-phase",
          phase: "discard",
        },
      ];
    }

    // Otherwise, the next player plays
    return [
      {
        if: finalRound && !state.finalRound,
        event: "final-round",
      },
      {
        if: state.phase !== "play",
        event: "change-phase",
        phase: "play",
      },
      {
        event: "change-player",
        playerId: nextPlayer(state),
      },
    ];
  },
};
