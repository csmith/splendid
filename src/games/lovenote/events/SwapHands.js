export default {
  name: "swap-hands",

  /**
   * @typedef HandToSwap
   * @type {object}
   * @property {string} id The ID of the player who will receive the hand.
   * @property {Object[]} hand The hand that will be received.
   */

  /**
   * @param players {HandToSwap[]} The details of the hands to swap.
   */
  create: function (players) {
    return {
      event: this.name,
      players,
    };
  },

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
