import _ from "lodash";
import {addObjects, subtractObjects} from "../../common/util.js";
import {findPlayer, nextPlayer} from "../../common/state.js";

const allowedTokens = ['emerald', 'ruby', 'diamond', 'sapphire', 'onyx'];

export default {
    name: 'take-tokens',

    available: function (state, {player}) {
        return player.id === state.turn;
    },

    perform: function (state, {player, tokens}) {
        const filteredTokens = Object.fromEntries(allowedTokens.map(t => [t, tokens[t] || 0]));
        const requestedTokens = _.sum(Object.values(filteredTokens));
        const doubleTokens = _.findKey(filteredTokens, amount => amount === 2);
        const totalAvailable = _.sum(allowedTokens.map(t => state.tokens[t]));

        if ((doubleTokens && requestedTokens !== 2) || (!doubleTokens && requestedTokens !== Math.min(3, totalAvailable))) {
            throw new Error(`You must draw 3 tokens of different types, or 2 of the same type`);
        }

        if (doubleTokens) {
            const potAmount = state.tokens[doubleTokens];
            if (potAmount < 4) {
                throw new Error(`Can't draw double ${doubleTokens} as only ${potAmount} are available`);
            }
        } else {
            Object.entries(filteredTokens).forEach(([type, amount]) => {
                if (amount > 1) {
                    throw new Error(`Can't draw more than 1 ${type} if drawing 3 total`);
                }

                const potAmount = state.tokens[type];
                if (potAmount < amount) {
                    throw new Error(`Can't draw ${type} as none are available`);
                }
            });
        }

        const playerData = findPlayer(state, player);
        const totalTokens = requestedTokens + _.sum(Object.values(playerData.tokens));
        const nextPhase = (totalTokens > 10) ? 'discard' : 'play';
        const nextTurn = (nextPhase === 'discard') ? state.turn : nextPlayer(state);

        return {
            ...state,
            phase: nextPhase,
            turn: nextTurn,
            players: {
                ...state.players,
                [player.id]: {
                    ...playerData,
                    tokens: addObjects(playerData.tokens, filteredTokens),
                }
            },
            tokens: subtractObjects(state.tokens, filteredTokens),
        }
    },
}