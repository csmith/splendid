export default {
  name: "round-over",

  /**
   * @param winningPlayerId {string} The ID of the player that won the round.
   */
  create: function (winningPlayerId) {
    return {
      event: this.name,
      winningPlayerId,
    };
  },

  perform: (state, { winningPlayerId }) => {
    state.players[winningPlayerId].points++;
  },
};
