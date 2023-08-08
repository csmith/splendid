import { addObjects, subtractObjects } from "../../../common/util.js";

export default {
  name: "take-tokens",

  perform: (state, { playerId, tokens }) => {
    state.tokens = subtractObjects(state.tokens, tokens);
    state.players[playerId].tokens = addObjects(state.players[playerId].tokens, tokens);
  },
};
