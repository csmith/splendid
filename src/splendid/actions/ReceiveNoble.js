import _ from "lodash";
import {canReceiveNoble} from "../util.js";

export default {
    name: 'receive-noble',

    available: function (state, {player}) {
        return player.id === state.turn;
    },

    perform: function (state, {player, noble}) {
        const index = _.findIndex(state.nobles, (n) => _.isEqual(n, noble));
        if (index === -1) {
            throw new Error('Noble not found');
        }

        if (!canReceiveNoble(state.players[player.id], noble)) {
            throw new Error('Not eligible for noble')
        }

        return _.concat(
            {
                ...state,
                nobles: _.concat(
                    _.slice(state.nobles, 0, index),
                    _.slice(state.nobles, index + 1)
                ),
                players: {
                    ...state.players,
                    [player.id]: {
                        ...state.players[player.id],
                        nobles: _.concat(state.players[player.id].nobles, noble),
                        points: state.players[player.id].points + 3,
                    }
                }
            },
            {
                action: 'end-turn',
            },
        )
    }
}