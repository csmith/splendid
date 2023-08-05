import _ from 'lodash';

export default {
    name: 'discard-reserve',

    perform: (state, {card}) => {
        state.players[state.turn].reserved = _.filter(
            state.players[state.turn].reserved,
            (c) => !_.isEqual(c, card)
        );
    }
}