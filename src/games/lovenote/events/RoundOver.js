export default {
  name: "round-over",

  /**
   * @param winningPlayerIds {string[]} The IDs of the players that won the round.
   */
  create: function (winningPlayerIds) {
    return {
      event: this.name,
      winningPlayerIds,
    };
  },

  perform: (state, { winningPlayerIds }) => {
    winningPlayerIds.forEach((playerId) => {
      state.players[playerId].points++;
    });
  },
};
