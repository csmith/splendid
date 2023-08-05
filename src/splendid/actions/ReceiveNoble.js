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
                event: 'receive-noble',
                noble,
            },
            {
                event: 'add-points',
                points: 3,
            },
            {
                action: 'end-turn',
            },
        )
    }
}