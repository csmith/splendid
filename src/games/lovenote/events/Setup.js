export default {
  name: "setup",

  perform: (state, { deck, unused, spare, tokensToWin }) => {
    state.deck = deck;
    state.unused = unused;
    state.spare = spare;
    state.tokensToWin = tokensToWin;
  },
};
