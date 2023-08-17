export default {
  name: "set-player-order",

  /**
   * @param order {string[]} Array of player IDs in the order they should play.
   */
  create: function (order) {
    return {
      event: this.name,
      order,
    };
  },

  perform: (state, { order }) => {
    state.turn = order[0];
    order.forEach((player, index) => {
      state.players[player].order = index;
    });
  },
};
