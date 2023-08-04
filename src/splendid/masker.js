import _ from 'lodash';

export default function (state, playerId) {
    return _.mapValues(state, (v, k) => {
        if (k === 'decks') {
            return v.map(deck => ({length: deck.length}));
        } else if (k === 'players') {
            return _.mapValues(v, (vv, player) => {
                if (player === playerId) {
                    return vv;
                } else {
                    return {
                        ...vv,
                        reserved: {length: vv.reserved.length},
                    }
                }
            });
        } else {
            return v;
        }
    })
}