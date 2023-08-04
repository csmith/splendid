import {replaceNth} from "../../common/util.js";

export default {
    name: 'deal',

    available: () => false,

    perform: function(state, {level, column}) {
        const card = state.decks[level][0];

        return {
            ...state,
            decks: replaceNth(state.decks, level, (d) => d.slice(1)),
            cards: replaceNth(state.cards, level, (row) => replaceNth(row, column, () => card)),
        }
    }
}