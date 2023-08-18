export default {
  name: "game-over",

  // TODO: Multiple players can win in the 2nd edition
  /**
   * @param winningPlayerId {string} The ID of the player that won the game.
   */
  create: function (winningPlayerId) {
    return {
      event: this.name,
      winningPlayerId,
    };
  },

  perform: () => {
    // Do nothing: informational only
  },
};
