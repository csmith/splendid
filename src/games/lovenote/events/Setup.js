export default {
  name: "setup",

  mask: function (playerId, data) {
    return {
      ...data,
      deck: data.deck.map((c) => ({ id: c.id })),
      unused: data.unused.map((c) => ({ id: c.id })),
    };
  },

  perform: (state, { deck, unused, spare, tokensToWin }) => {
    state.deck = deck;
    state.unused = unused;
    state.spare = spare;
    state.tokensToWin = tokensToWin;
  },
};
