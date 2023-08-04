import {findPlayer} from "../../common/state.js";

export default {
    name: 'join',

    available: function (state, {player}) {
        return findPlayer(state, player) === undefined;
    },

    perform: function (state, {player}) {
        return {
            ...state,
            players: {
                ...state.players,
                [player.id]: {
                    details: player,
                    reserved: [],
                    nobles: [],
                    order: undefined,
                    points: 0,
                    tokens: {
                        emerald: 0,
                        diamond: 0,
                        sapphire: 0,
                        onyx: 0,
                        ruby: 0,
                        gold: 0,
                    },
                    bonuses: {
                        emerald: 0,
                        diamond: 0,
                        sapphire: 0,
                        onyx: 0,
                        ruby: 0,
                    },
                }
            }
        }
    }
}