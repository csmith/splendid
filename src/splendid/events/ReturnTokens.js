import {addObjects, subtractObjects} from "../../common/util.js";

export default {
    name: 'return-tokens',

    perform: (state, {tokens}) => {
        state.tokens = addObjects(state.tokens, tokens);
        state.players[state.turn].tokens = subtractObjects(state.players[state.turn].tokens, tokens);
    }
}