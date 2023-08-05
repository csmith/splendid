import _ from "lodash";
import {addObjects, subtractObjects} from "../../common/util.js";
import {findPlayer} from "../../common/state.js";

export default {
    name: 'reserve-card',

    available: function (state, {player}) {
        return player.id === state.turn
            && state.players[player.id].reserved.length < 3;
    },

    perform: function (state, {card}) {
        const index = _.findIndex(state.cards[card.level - 1], (c) => _.isEqual(c, card));
        if (index === -1) {
            throw new Error('Card not found');
        }

        const getsGold = state.tokens.gold > 0;

        return _.concat(
            {
                event: 'discard-card',
                card,
            },
            {
                event: 'reserve-card',
                card,
            },
            {
                if: getsGold,
                event: 'take-tokens',
                tokens: {gold: 1},
            },
            {
                action: 'deal',
                args: {
                    level: card.level,
                },
            },
            {
                action: 'end-turn',
            },
        );
    }
}