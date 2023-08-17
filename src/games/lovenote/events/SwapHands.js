export default {
  name: "swap-hands",

  mask: function (playerId, data) {
    return {
      ...data,
      players: data.players.map((p) => ({
        ...p,
        hand: p.id === playerId ? p.hand : p.hand.map((c) => ({ id: c.id })),
      })),
    };
  },

  perform: (state, { players }) => {
    state.players[players[0].id].hand = players[0].hand;
    state.players[players[1].id].hand = players[1].hand;
  },
};
