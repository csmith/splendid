import _ from 'lodash';

export default {
    name: 'discard-card',

    perform: (state, {card}) => {
        state.cards = state.cards.map(row => row.map((c) => _.isEqual(c, card) ? undefined : c));
    }
}