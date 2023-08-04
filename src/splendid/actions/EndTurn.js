import {isLastPlayer, nextPlayer} from "../../common/state.js";
import _ from "lodash";
import {canReceiveNoble} from "../util.js";

export default {
    name: 'end-turn',

    available: () => false,

    perform: (state) => {
        const currentPlayer = state.players[state.turn];
        const finalRound = state.finalRound || Object.values(state.players).some((p) => p.points >= 15);

        // First, if there are any nobles that can be received, the player must receive one
        if (state.phase === 'play' && state.nobles.some((n) => canReceiveNoble(currentPlayer, n))) {
            return {
                ...state,
                finalRound,
                phase: 'noble'
            }
        }

        // If this is the final round and this is the last player, the game is over
        if (finalRound && isLastPlayer(state, currentPlayer.details)) {
            return {
                ...state,
                finalRound,
                phase: 'end',
                turn: undefined,
            }
        }

        // If we're not finishing immediately, and the current player is over their token limit, they must discard
        if (_.sum(Object.values(currentPlayer.tokens)) > 10) {
            return {
                ...state,
                finalRound,
                phase: 'discard'
            }
        }

        // Otherwise, the next player plays
        return {
            ...state,
            finalRound,
            phase: 'play',
            turn: nextPlayer(state),
        }
    }
}