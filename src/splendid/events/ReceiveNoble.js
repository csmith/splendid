import _ from 'lodash';

export default {
    name: 'receive-noble',

    perform: (state, {noble}) => {
        state.nobles = _.filter(state.nobles, (n) => !_.isEqual(n, noble));
        state.players[state.turn].nobles = [...state.players[state.turn].nobles, noble];
    }
}