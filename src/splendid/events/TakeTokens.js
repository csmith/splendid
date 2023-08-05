import {addObjects, subtractObjects} from "../../common/util.js";

export default {
    name: 'take-tokens',

    perform: (state, {tokens}) => {
        state.tokens = subtractObjects(state.tokens, tokens);
        state.players[state.turn].tokens = addObjects(state.players[state.turn].tokens, tokens);
    }
}