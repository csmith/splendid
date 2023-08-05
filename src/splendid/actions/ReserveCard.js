import _ from "lodash";
import {addObjects, subtractObjects} from "../../common/util.js";
import {findPlayer} from "../../common/state.js";

export default {
    name: 'reserve-card',

    available: function (state, {player}) {
        return player.id === state.turn
            && state.players[player.id].reserved.length < 3;
    },

    perform: function (state, {player, card}) {
        const index = _.findIndex(state.cards[card.level - 1], (c) => _.isEqual(c, card));
        if (index === -1) {
            throw new Error('Card not found');
        }

        const playerData = findPlayer(state, player);
        const getsGold = state.tokens.gold > 0;

        return _.concat(
            {
                ...state,
                players: {
                    ...state.players,
                    [player.id]: {
                        ...playerData,
                        tokens: addObjects(playerData.tokens, {gold: getsGold ? 1 : 0}),
                        reserved: _.concat(playerData.reserved, card),
                    }
                },
                tokens: subtractObjects(state.tokens, {gold: getsGold ? 1 : 0}),
            },
            {
                event: 'discard-card',
                card,
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