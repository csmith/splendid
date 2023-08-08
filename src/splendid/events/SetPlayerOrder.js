export default {
  name: "set-player-order",

  perform: (state, { order }) => {
    state.turn = order[0];
    order.forEach((player, index) => {
      state.players[player].order = index;
    });
  },
};
