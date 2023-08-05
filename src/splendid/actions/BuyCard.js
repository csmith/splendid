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

        return [
            {
                if: _.sum(Object.values(deductions)) > 0,
                event: 'return-tokens',
                tokens: deductions,
            },
            {
                if: reserveIndex !== -1,
                event: 'discard-reserve',
                card,
            },
            {
                if: index !== -1,
                event: 'discard-card',
                card,
            },
            {
                if: card.points > 0,
                event: 'add-points',
                points: card.points,
            },
            {
                event: 'add-bonus',
                type: card.bonus,
            },
            {
                if: index > -1,
                action: 'deal',
                args: {
                    level: card.level - 1,
                    column: index,
                },
            },
            {
                action: 'end-turn',
            },
        ];
    },
}