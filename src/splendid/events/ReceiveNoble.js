import _ from 'lodash';

export default {
    name: 'receive-noble',

    perform: (state, {playerId, noble}) => {
        state.nobles = _.filter(state.nobles, (n) => !_.isEqual(n, noble));
        state.players[playerId].nobles = [...state.players[playerId].nobles, noble];
    }
}