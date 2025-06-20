export default {
  name: "game-over",

  /**
   * @param winningPlayerIds {string[]} The IDs of the players that won the game.
   */
  create: function (winningPlayerIds) {
    return {
      event: this.name,
      winningPlayerIds,
    };
  },

  perform: () => {
    // Do nothing: informational only
  },
};
