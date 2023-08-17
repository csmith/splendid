import { addObjects, subtractObjects } from "../../../common/util.js";

export default {
  name: "take-tokens",

  /**
   * @param playerId {string} the player who is taking the tokens
   * @param tokens {Object.<string, number>} the tokens to take
   */
  create: function (playerId, tokens) {
    return {
      event: this.name,
      playerId,
      tokens,
    };
  },

  perform: (state, { playerId, tokens }) => {
    state.tokens = subtractObjects(state.tokens, tokens);
    state.players[playerId].tokens = addObjects(state.players[playerId].tokens, tokens);
  },
};
