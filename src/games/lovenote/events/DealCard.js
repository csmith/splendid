export default {
  name: "deal-card",

  perform: (state, { playerId, card }) => {
    state.deck = state.deck.slice(1);
    state.players[playerId].hand.push(card);
  },
};
