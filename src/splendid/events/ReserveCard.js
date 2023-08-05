import _ from 'lodash';

export default {
    name: 'reserve-card',

    perform: (state, {card}) => {
        state.players[state.turn].reserved = _.concat(state.players[state.turn].reserved, card);
    }
}