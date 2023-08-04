import _ from "lodash";
import {addObjects, subtractObjects} from "../../common/util.js";
import {findPlayer, isLastPlayer, nextPlayer} from "../../common/state.js";

const allowedTokens = ['emerald', 'ruby', 'diamond', 'sapphire', 'onyx', 'gold'];

export default {
    name: 'discard-tokens',

    available: function (state, {player}) {
        return player.id === state.turn;
    },

    perform: function (state, {player, tokens}) {
        const filteredTokens = Object.fromEntries(allowedTokens.map(t => [t, tokens[t] || 0]));
        const playerData = findPlayer(state, player);
        const totalTokens = _.sum(Object.values(playerData.tokens));
        const discardedTokens = _.sum(Object.values(filteredTokens));
        if (totalTokens - discardedTokens !== 10) {
            throw new Error(`You must discard ${totalTokens - 10} tokens`);
        }

        Object.entries(playerData.tokens).forEach(([token, count]) => {
            if (filteredTokens[token] > count) {
                throw new Error(`You only have ${count} ${token} tokens`);
            }
        });

        return _.concat(
            {
                ...state,
                players: {
                    ...state.players,
                    [player.id]: {
                        ...playerData,
                        tokens: subtractObjects(playerData.tokens, filteredTokens),
                    }
                },
                tokens: addObjects(state.tokens, filteredTokens),
            },
            {
                action: 'end-turn',
            },
        );
    },
}