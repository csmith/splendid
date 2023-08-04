import _ from "lodash";
import {addObjects, subtractObjects} from "../../common/util.js";
import {findPlayer} from "../../common/state.js";
import {canAffordCard, costForCard} from "../util.js";

export default {
    name: 'buy-card',

    available: function (state, {player}) {
        return player.id === state.turn;
    },

    perform: function (state, {player, card}) {
        const index = _.findIndex(state.cards[card.level - 1], (c) => _.isEqual(c, card));
        const reserveIndex = _.findIndex(state.players[player.id].reserved, (c) => _.isEqual(c, card));

        if (index === -1 && reserveIndex === -1) {
            throw new Error('Card not found');
        }

        const playerData = findPlayer(state, player);
        if (!canAffordCard(playerData, card)) {
            throw new Error('Cannot afford card');
        }

        const cost = costForCard(playerData, card);
        const remaining = subtractObjects(playerData.tokens, cost);
        const missing = _.mapValues(Object.fromEntries(Object.entries(remaining).filter((a) => a[1] < 0)), (a) => -a);
        const missingCount = _.sum(Object.values(missing));

        const deductions = {...subtractObjects(cost, missing), gold: missingCount};
        const newScore = playerData.points + card.points;

        return _.concat(
            {
                ...state,
                players: {
                    ...state.players,
                    [player.id]: {
                        ...playerData,
                        points: newScore,
                        tokens: subtractObjects(playerData.tokens, deductions),
                        reserved: reserveIndex === -1 ? playerData.reserved : [
                            ...playerData.reserved.slice(0, reserveIndex),
                            ...playerData.reserved.slice(reserveIndex + 1),
                        ],
                        bonuses: {
                            ...playerData.bonuses,
                            [card.bonus]: playerData.bonuses[card.bonus] + 1,
                        }
                    }
                },
                tokens: addObjects(state.tokens, deductions),
            },
            index === -1 ? [] : {
                action: 'deal',
                args: {
                    level: card.level - 1,
                    column: index,
                },
            },
            {
                action: 'end-turn',
            },
        );
    },
}