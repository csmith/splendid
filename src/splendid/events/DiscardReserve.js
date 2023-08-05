import _ from 'lodash';

export default {
    name: 'discard-reserve',

    perform: (state, {playerId, card}) => {
        state.players[playerId].reserved = _.filter(
            state.players[playerId].reserved,
            (c) => !_.isEqual(c, card)
        );
    }
}