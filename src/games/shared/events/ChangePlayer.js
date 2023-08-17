export default {
  name: "change-player",

  /**
   * @param playerId {string|undefined} The ID of the player whose turn is next
   */
  create: function (playerId) {
    return {
      event: this.name,
      playerId,
    };
  },

  perform: function (state, { playerId }) {
    state.turn = playerId;
  },
};
