// TODO: Maybe add a ready check for everyone.
import {countPlayers, findPlayer} from "../../common/state.js";
import cards from "../data/cards.js";
import nobles from "../data/nobles.js";
import _ from "lodash";

const tokensToRemovePerPlayerCount = {
    2: 3,
    3: 2,
    4: 0,
}

export default {
    name: 'start',

    available: function (state, {player}) {
        const count = countPlayers(state);
        const isPlayer = findPlayer(state, player);
        return isPlayer && count >= 2 && count <= 4;
    },

    perform: function (state) {
        const players = countPlayers(state);
        const tokensToRemove = tokensToRemovePerPlayerCount[players];
        const decks = _.times(3, (level) => _.filter(cards, (c) => c.level === level + 1));
        const turnOrder = _.shuffle(Object.keys(state.players));

        return _.concat(
            // Sort the tokens out, shuffle the decks, and assign turn order.
            {
                ...state,
                tokens: {
                    emerald: 7 - tokensToRemove,
                    diamond: 7 - tokensToRemove,
                    sapphire: 7 - tokensToRemove,
                    onyx: 7 - tokensToRemove,
                    ruby: 7 - tokensToRemove,
                    gold: 5,
                },
                nobles: _.take(_.shuffle(nobles), players + 1),
                decks: decks.map((d) => _.shuffle(d)),
                players: _.mapValues(state.players, (p, id) => ({
                    ...p,
                    order: turnOrder.indexOf(id),
                })),
                turn: turnOrder[0],
            },
            // Deal four cards from each deck.
            _.flatMap(decks, (d, level) =>
                _.times(4, (column) => ({
                    action: 'deal',
                    args: {level, column},
                }))
            ),
            // Begin the play phase
            {
                action: 'change-phase',
                args: {phase: 'play'},
            }
        )
    }
}