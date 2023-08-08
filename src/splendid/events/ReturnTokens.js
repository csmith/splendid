import { addObjects, subtractObjects } from "../../common/util.js";

export default {
  name: "return-tokens",

  perform: (state, { playerId, tokens }) => {
    state.tokens = addObjects(state.tokens, tokens);
    state.players[playerId].tokens = subtractObjects(state.players[playerId].tokens, tokens);
  },
};
