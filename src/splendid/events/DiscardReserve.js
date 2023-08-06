import _ from 'lodash';

export default {
    name: 'discard-reserve',

    perform: (state, {playerId, card}) => {
        _.remove(state.players[playerId].reserved, (c) => c.id === card.id);
    }
}