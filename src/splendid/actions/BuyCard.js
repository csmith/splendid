import _ from "lodash";
import {addObjects, subtractObjects} from "../../common/util.js";
import {findPlayer, isLastPlayer, nextPlayer} from "../../common/state.js";
import {canAffordCard, canReceiveNoble, costForCard} from "../util.js";

export default {
    name: 'buy-card',

    available: function (state, {player}) {
        return player.id === state.turn;
    },

    perform: function (state, {player, card}) {
        const index = _.findIndex(state.cards[card.level - 1], (c) => _.isEqual(c, card));
        if (index === -1) {
            // TODO: it could be a reserve card
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

        const canReceiveNobles = state.nobles.some((noble) => canReceiveNoble(playerData, noble));
        const newScore = playerData.points + card.points;
        const nextTurn = canReceiveNobles ? state.turn : nextPlayer(state);
        const finalRound = state.finalRound || newScore >= 15;
        const nextPhase = canReceiveNobles ? 'noble' : (finalRound && isLastPlayer(state, player) ? 'end' : 'play');

        return _.concat(
            {
                ...state,
                turn: nextTurn,
                phase: nextPhase,
                finalRound: finalRound,
                players: {
                    ...state.players,
                    [player.id]: {
                        ...playerData,
                        points: newScore,
                        tokens: subtractObjects(playerData.tokens, deductions),
                        bonuses: {
                            ...playerData.bonuses,
                            [card.bonus]: playerData.bonuses[card.bonus] + 1,
                        }
                    }
                },
                tokens: addObjects(state.tokens, deductions),
            },
            {
                action: 'deal',
                args: {
                    level: card.level - 1,
                    column: index,
                },
            }
        );
    },
}