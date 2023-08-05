import _ from 'lodash';

export default {
    name: 'place-card',

    perform: (state, {card}) => {
        const index = _.findIndex(state.cards[card.level-1], (c) => c === undefined);
        state.cards[card.level-1][index] = card;
    }
}