import _ from "lodash";
import {addObjects, replaceNth, subtractObjects} from "../../common/util.js";
import {findPlayer} from "../../common/state.js";

export default {
    name: 'reserve-card-from-deck',

    available: function (state, {player}) {
        return player.id === state.turn
            && state.players[player.id].reserved.length < 3;
    },

    perform: function (state, {player, level}) {
        if (level < 1 || level > 3) {
            throw new Error('Invalid level');
        }

        if (state.decks[level-1].length === 0) {
            throw new Error('Deck is empty');
        }

        const card = state.decks[level-1][0];
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
                decks: replaceNth(state.decks, level-1, (d) => d.slice(1)),
                tokens: subtractObjects(state.tokens, {gold: getsGold ? 1 : 0}),
            },
            {
                action: 'end-turn',
            },
        );
    }
}