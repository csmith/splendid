export default {
  name: "eliminate-player",

  /**
   * @param playerId {string} The ID of the player being eliminated
   * @param reason {string} The reason the player is being eliminated
   */
  create: function (playerId, reason) {
    return {
      event: this.name,
      playerId,
      reason,
    };
  },

  perform: (state, { playerId }) => {
    state.players[playerId].eliminated = true;
  },
};
