import { addObjects, subtractObjects } from "../../../common/util.js";

export default {
  name: "return-tokens",

  /**
   * @param playerId {string} The player who is returning tokens
   * @param tokens {Object.<string,number>} The tokens to return
   */
  create: function (playerId, tokens) {
    return {
      event: this.name,
      playerId,
      tokens,
    };
  },

  perform: (state, { playerId, tokens }) => {
    state.tokens = addObjects(state.tokens, tokens);
    state.players[playerId].tokens = subtractObjects(state.players[playerId].tokens, tokens);
  },
};
