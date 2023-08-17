export default {
  name: "setup",

  /**
   * @param tokens {Object.<string, number>} The initial pool of tokens
   * @param nobles {Object[]} The available nobles
   * @param decks {Object[][]} The three decks of cards
   * @returns {{nobles, tokens, decks, event: string}}
   */
  create: function (tokens, nobles, decks) {
    return {
      event: this.name,
      tokens,
      nobles,
      decks,
    };
  },

  perform: (state, { tokens, nobles, decks }) => {
    state.tokens = tokens;
    state.nobles = nobles;
    state.decks = decks;
  },
};
