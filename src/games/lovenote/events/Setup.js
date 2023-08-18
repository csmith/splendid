export default {
  name: "setup",

  /**
   * @param deck {Object[]} The randomly sorted deck.
   * @param unused {Object[]} The cards removed from the deck for two player games.
   * @param spare {Object[]} The card removed from the deck for use in the final round.
   * @param tokensToWin {number} The number of tokens required to win.
   */
  create: function (deck, unused, spare, tokensToWin) {
    return {
      event: this.name,
      deck,
      unused,
      spare,
      tokensToWin,
    };
  },

  mask: function (playerId, data) {
    return {
      ...data,
      deck: data.deck.map((c) => ({ id: c.id })),
      spare: data.spare.map((c) => ({ id: c.id })),
    };
  },

  perform: (state, { deck, unused, spare, tokensToWin }) => {
    state.deck = deck;
    state.unused = unused;
    state.spare = spare;
    state.tokensToWin = tokensToWin;
  },
};
