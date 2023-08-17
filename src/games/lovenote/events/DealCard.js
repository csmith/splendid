export default {
  name: "deal-card",

  mask: function (playerId, data) {
    if (playerId === data.playerId) {
      return data;
    }

    return {
      ...data,
      hand: data.hand.map((c) => ({ id: c.id })),
    };
  },

  perform: (state, { playerId, card }) => {
    state.deck = state.deck.slice(1);
    state.players[playerId].hand.push(card);
  },
};
