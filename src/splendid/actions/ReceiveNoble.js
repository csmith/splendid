import _ from "lodash";

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
                        nobles: _.concat(state.players[player.id].nobles, noble)
                    }
                }
            },
            {
                action: 'end-turn',
            },
        )
    }
}